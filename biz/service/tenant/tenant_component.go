package tenant

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/k8s"
)

func (a *TenantService) AddTenantComponent(ctx context.Context, req *freezonex_openiiot_api.AddTenantComponentRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddTenantComponentResponse, error) {

	_, tenantName, err := a.CheckTenant(ctx, req.TenantId, req.TenantName)
	if err != nil {
		logs.Error(ctx, "event=AddTenantComponent error=%v", err.Error())
		return nil, err
	}
	if req.ComponentName != "nodered" && req.ComponentName != "grafana" {
		return nil, errors.New("Only nodered and grafana are allowed to have multiple instances in one tenant")
	}

	ctx = context.WithValue(ctx, "tenantName", tenantName)

	k8sUns := k8s.K8sUns{TenantName: tenantName, ComponentName: req.ComponentName}

	number, err := a.k8s.GetNextAvailableComponentDeploymentNumber(ctx, k8sUns)
	if err != nil {
		logs.Error(ctx, "event=AddTenantComponent error=%v", err.Error())
		return nil, err
	}

	k8sUns.Number = strconv.Itoa(number)

	switch req.ComponentName {
	case "nodered":
		if err := a.k8s.CreateNoderedComponent(ctx, k8sUns); err != nil {
			return nil, err
		}
	case "grafana":
		if err := a.k8s.CreateGrafanaComponent(ctx, k8sUns); err != nil {
			return nil, err
		}
	}

	resp := new(freezonex_openiiot_api.AddTenantComponentResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.ComponentUri = a.k8s.GetComponentRootUrl(k8sUns) + "/"

	return resp, nil
}

func (a *TenantService) DeleteTenantComponent(ctx context.Context, req *freezonex_openiiot_api.DeleteTenantComponentRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteTenantComponentResponse, error) {

	_, tenantName, err := a.CheckTenant(ctx, req.TenantId, req.TenantName)
	if err != nil {
		logs.Error(ctx, "event=DeleteTenantComponent error=%v", err.Error())
		return nil, err
	}
	ctx = context.WithValue(ctx, "tenantName", tenantName)

	k8sUns := k8s.K8sUns{TenantName: tenantName}
	deploymentNames, err := a.k8s.GetRuntimeDeploymentNames(ctx, k8sUns)
	if err != nil {
		return nil, fmt.Errorf("failed to list deployments: %w", err)
	}

	if !a.FindComponentName(deploymentNames, req.ComponentName) {
		return nil, fmt.Errorf("cannot find component %s in tenant %s", req.ComponentName, req.TenantName)
	}

	componentName, number, err := a.SplitComponentNameStrict(req.ComponentName)
	if err != nil {
		return nil, err
	}

	k8sUns = k8s.K8sUns{TenantName: tenantName, ComponentName: componentName, Number: number}
	if err := a.k8s.DeleteComponent(ctx, k8sUns); err != nil {
		return nil, err
	}

	resp := new(freezonex_openiiot_api.DeleteTenantComponentResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *TenantService) GetTenantComponent(ctx context.Context, req *freezonex_openiiot_api.GetTenantComponentRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetTenantComponentResponse, error) {

	_, tenantName, err := a.CheckTenant(ctx, req.TenantId, req.TenantName)
	if err != nil {
		logs.Error(ctx, "event=GetTenantComponent error=%v", err.Error())
		return nil, err
	}
	ctx = context.WithValue(ctx, "tenantName", tenantName)

	componentName, number := a.SplitComponentNameLoose(req.ComponentName)
	k8sUns := k8s.K8sUns{TenantName: tenantName, ComponentName: componentName, Number: number}

	pods, err := a.k8s.GetRuntimePods(ctx, k8sUns, false)
	if err != nil {
		logs.Error(ctx, "event=GetTenantComponent error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetTenantComponentResponse)
	data := make([]*freezonex_openiiot_api.TenantComponent, 0)
	for _, pod := range pods {
		re := regexp.MustCompile(`^(.*)-[^-]+-[^-]+$`)
		deploymentName := re.ReplaceAllString(pod.Name, `$1`) //extrace deploymentName from openiiot-nodered1-6bcf7f7ffc-nx44x

		// Split the string by dash
		parts := strings.Split(deploymentName, "-")

		// Get the last element from the parts slice
		componentName := parts[len(parts)-1]

		uri := "/" + tenantName + "/" + componentName + "/"

		// Calculate Ready containers
		readyContainers := 0
		for _, containerStatus := range pod.Status.ContainerStatuses {
			if containerStatus.Ready {
				readyContainers++
			}
		}
		ready := fmt.Sprintf("%d/%d", readyContainers, len(pod.Status.ContainerStatuses))

		// Get Pod status
		status := string(pod.Status.Phase)

		// Calculate container restarts
		restarts := int32(0)
		for _, containerStatus := range pod.Status.ContainerStatuses {
			restarts += containerStatus.RestartCount
		}

		// Calculate Pod age
		age := time.Since(pod.CreationTimestamp.Time).Round(time.Second).String()

		// Retrieve the alias label from pod metadata
		alias := pod.Labels["alias"]
		if alias == "" {
			alias = componentName // Default value if alias label is not present
		}
		fmt.Printf("Pod alias: %s\n", alias)

		data = append(data, &freezonex_openiiot_api.TenantComponent{
			Name:     deploymentName,
			Uri:      uri,
			Ready:    ready,
			Status:   status,
			Restarts: strconv.Itoa(int(restarts)),
			Age:      age,
			Alias:    alias,
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *TenantService) UpdateTenantComponent(ctx context.Context, req *freezonex_openiiot_api.UpdateTenantComponentRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateTenantComponentResponse, error) {

	_, tenantName, err := a.CheckTenant(ctx, req.TenantId, req.TenantName)
	if err != nil {
		logs.Error(ctx, "event=UpdateTenantComponent error=%v", err.Error())
		return nil, err
	}
	ctx = context.WithValue(ctx, "tenantName", tenantName)

	k8sUns := k8s.K8sUns{TenantName: tenantName}
	deploymentNames, err := a.k8s.GetRuntimeDeploymentNames(ctx, k8sUns)
	if err != nil {
		return nil, fmt.Errorf("failed to list deployments: %w", err)
	}

	if !a.FindComponentName(deploymentNames, req.ComponentName) {
		return nil, fmt.Errorf("cannot find component %s in tenant %s", req.ComponentName, req.TenantName)
	}

	componentName, number, err := a.SplitComponentNameStrict(req.ComponentName)
	if err != nil {
		return nil, err
	}

	k8sUns = k8s.K8sUns{TenantName: tenantName, ComponentName: componentName, Number: number, Alias: req.ComponentAlias}
	if err := a.k8s.UpdateComponent(ctx, k8sUns); err != nil {
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateTenantComponentResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
