package application

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/k8s"
)

func (a *ApplicationService) AddApplication(ctx context.Context, req *freezonex_openiiot_api.AddApplicationRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddApplicationResponse, error) {
	_, tenantName, err := a.tenant.CheckTenant(ctx, req.TenantId, req.TenantName)
	if err != nil {
		logs.Error(ctx, "event=AddApplication error=%v", err.Error())
		return nil, err
	}

	// check application name if meet requirement
	err = a.CheckApplicationName(req.ApplicationName)
	if err != nil {
		logs.Error(ctx, "event=AddApplication error=%v", err.Error())
		return nil, err
	}

	// check if application exist
	exist, err := a.ExistApplication(ctx, tenantName, req.ApplicationName)
	if err != nil {
		logs.Error(ctx, "event=AddApplication error=%v", err.Error())
		return nil, err
	}
	if exist {
		return nil, fmt.Errorf("application name already exist: %s", req.ApplicationName)
	}

	// Multipart form
	form, _ := c.MultipartForm()
	fileHeaders := form.File["source_file"]

	// Build docker image
	imageName, err := a.BuildDockerImage(req.ApplicationName, req.ComponentType, fileHeaders)
	if err != nil {
		return nil, fmt.Errorf("failed to build Docker image: %v", err)
	}

	// Tag and push docker image to resistry
	err = a.TagAndPushDockerImage(imageName, "registry:5000")
	if err != nil {
		return nil, fmt.Errorf("failed to tag and push Docker image: %v", err)
	}

	ctx = context.WithValue(ctx, "tenantName", tenantName)

	k8sUns := k8s.K8sUns{TenantName: tenantName, DeploymentCategory: "app", ComponentName: req.ApplicationName, ComponentType: req.ComponentType}

	switch req.ComponentType {
	case "frontend":
		if err := a.k8s.CreateApplicationFrontendComponent(ctx, k8sUns); err != nil {
			return nil, err
		}
	case "backend":
		if err := a.k8s.CreateApplicationBackendComponent(ctx, k8sUns); err != nil {
			return nil, err
		}
	case "db":
		if err := a.k8s.CreateApplicationDbComponent(ctx, k8sUns); err != nil {
			return nil, err
		}
	}

	resp := new(freezonex_openiiot_api.AddApplicationResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.ApplicationUri = a.k8s.GetComponentRootUrl(k8sUns) + "/"

	return resp, nil
}

// GetApplication will get application record in condition
func (a *ApplicationService) GetApplication(ctx context.Context, req *freezonex_openiiot_api.GetApplicationRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetApplicationResponse, error) {
	_, tenantName, err := a.tenant.CheckTenant(ctx, req.TenantId, req.TenantName)
	if err != nil {
		logs.Error(ctx, "event=GetApplication error=%v", err.Error())
		return nil, err
	}
	ctx = context.WithValue(ctx, "tenantName", tenantName)

	k8sUns := k8s.K8sUns{TenantName: tenantName, DeploymentCategory: "app", ComponentName: req.ApplicationName, ComponentType: req.ComponentType}

	deployments, err := a.k8s.GetDeploymentsByFuzzyName(ctx, k8sUns)
	if err != nil {
		logs.Error(ctx, "event=GetApplication error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetApplicationResponse)
	data := make([]*freezonex_openiiot_api.Application, 0)
	componentMap := make(map[string][]*freezonex_openiiot_api.ApplicationComponent)

	for _, deployment := range deployments {
		if !strings.HasPrefix(deployment.Name, "openiiot-app") {
			continue
		}

		re := regexp.MustCompile(`^openiiot-app-([a-z0-9-]+)-(frontend|backend|db)$`)
		match := re.FindStringSubmatch(deployment.Name)
		if match == nil {
			return nil, fmt.Errorf("deployment name does not match the expected pattern")
		}
		applicationName := match[1]
		componentType := match[2]

		uri := "/" + tenantName + "/" + applicationName + "/"

		// Calculate Pod age
		age := time.Since(deployment.CreationTimestamp.Time).Round(time.Second).String()

		// Retrieve the alias label from pod metadata
		alias := deployment.Labels["alias"]
		if alias == "" {
			alias = applicationName // Default value if alias label is not present
		}

		// Extract Ready: number of ready replicas vs desired replicas (e.g., 1/1)
		ready := fmt.Sprintf("%d/%d", deployment.Status.ReadyReplicas, *deployment.Spec.Replicas)

		// Extract Status: if AVAILABLE >= 1, status is Running, else Stopped
		status := "Stopped"
		if deployment.Status.AvailableReplicas >= 1 {
			status = "Running"
		}

		// Extract Revision: from annotations (deployment.kubernetes.io/revision)
		revision := deployment.Annotations["deployment.kubernetes.io/revision"]
		if revision == "" {
			revision = "1" // Default revision if annotation not present
		}

		// Extract CreationTimestamp
		createTime := deployment.CreationTimestamp.Time.Format("2006-01-02 15:04:05 -0700")

		applicationComponent := &freezonex_openiiot_api.ApplicationComponent{
			ComponentType: componentType,
			Ready:         ready,
			Status:        status,
			Revision:      revision,
			CreateTime:    createTime,
			Age:           age,
		}

		// If components exist, append; otherwise initialize a new entry
		components := componentMap[applicationName]
		components = append(components, applicationComponent)
		componentMap[applicationName] = components

		// Only add to response data once
		if len(components) == 1 {
			data = append(data, &freezonex_openiiot_api.Application{
				Name:       applicationName,
				Uri:        uri,
				Components: components,
			})
		}
	}

	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// GetApplication will get application record in condition
func (a *ApplicationService) ExistApplication(ctx context.Context, tenantName string, applicationName string) (bool, error) {

	k8sUns := k8s.K8sUns{TenantName: tenantName, DeploymentCategory: "app", ComponentName: applicationName}

	deployments, err := a.k8s.GetDeploymentsByFuzzyName(ctx, k8sUns)
	if err != nil {
		logs.Error(ctx, "event=ExistApplication error=%v", err.Error())
		return false, err
	}

	for _, deployment := range deployments {
		if !strings.HasPrefix(deployment.Name, "openiiot-app") {
			continue
		}

		re := regexp.MustCompile(`^openiiot-app-([a-z0-9-]+)-(frontend|backend|db)$`)
		match := re.FindStringSubmatch(deployment.Name)
		if match == nil {
			return false, fmt.Errorf("deployment name does not match the expected pattern")
		}

		if applicationName == match[1] {
			return true, nil
		}
	}

	return false, nil
}

// StartApplication will start application, if component_type is empty, then it will start all component of this application
func (a *ApplicationService) StartApplication(ctx context.Context, req *freezonex_openiiot_api.StartApplicationRequest, c *app.RequestContext) (*freezonex_openiiot_api.StartApplicationResponse, error) {

	_, tenantName, err := a.tenant.CheckTenant(ctx, req.TenantId, req.TenantName)
	if err != nil {
		logs.Error(ctx, "event=StartApplication error=%v", err.Error())
		return nil, err
	}
	ctx = context.WithValue(ctx, "tenantName", tenantName)

	// check if application exist
	exist, err := a.ExistApplication(ctx, tenantName, req.ApplicationName)
	if err != nil {
		logs.Error(ctx, "event=Startpplication error=%v", err.Error())
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("application name does not exist: %s", req.ApplicationName)
	}

	k8sUns := k8s.K8sUns{TenantName: tenantName, DeploymentCategory: "app", ComponentName: req.ApplicationName, ComponentType: req.ComponentType}

	deployments, err := a.k8s.GetDeploymentsByFuzzyName(ctx, k8sUns)
	if err != nil {
		logs.Error(ctx, "event=StartApplication error=%v", err.Error())
		return nil, err
	}

	for _, deployment := range deployments {
		/*if !strings.HasPrefix(deployment.Name, "openiiot-app-"+req.ApplicationName) {
			continue
		}*/
		re := regexp.MustCompile(`^openiiot-app-([a-z0-9-]+)-(frontend|backend|db)$`)
		match := re.FindStringSubmatch(deployment.Name)
		if match == nil {
			return nil, fmt.Errorf("deployment name does not match the expected pattern")
		}
		k8sUns.ComponentType = match[2]

		a.k8s.StartDeployment(ctx, k8sUns)
	}

	resp := new(freezonex_openiiot_api.StartApplicationResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// StopApplication will stop application, if component_type is empty, then it will stop all component of this application
func (a *ApplicationService) StopApplication(ctx context.Context, req *freezonex_openiiot_api.StopApplicationRequest, c *app.RequestContext) (*freezonex_openiiot_api.StopApplicationResponse, error) {

	_, tenantName, err := a.tenant.CheckTenant(ctx, req.TenantId, req.TenantName)
	if err != nil {
		logs.Error(ctx, "event=StopApplication error=%v", err.Error())
		return nil, err
	}
	ctx = context.WithValue(ctx, "tenantName", tenantName)

	// check if application exist
	exist, err := a.ExistApplication(ctx, tenantName, req.ApplicationName)
	if err != nil {
		logs.Error(ctx, "event=StopApplication error=%v", err.Error())
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("application name already not exist: %s", req.ApplicationName)
	}

	k8sUns := k8s.K8sUns{TenantName: tenantName, DeploymentCategory: "app", ComponentName: req.ApplicationName, ComponentType: req.ComponentType}

	deployments, err := a.k8s.GetDeploymentsByFuzzyName(ctx, k8sUns)
	if err != nil {
		logs.Error(ctx, "event=StopApplication error=%v", err.Error())
		return nil, err
	}

	for _, deployment := range deployments {
		/*if !strings.HasPrefix(deployment.Name, "openiiot-app-"+req.ApplicationName) {
			continue
		}*/
		re := regexp.MustCompile(`^openiiot-app-([a-z0-9-]+)-(frontend|backend|db)$`)
		match := re.FindStringSubmatch(deployment.Name)
		if match == nil {
			return nil, fmt.Errorf("deployment name does not match the expected pattern")
		}
		k8sUns.ComponentType = match[2]

		a.k8s.StopDeployment(ctx, k8sUns)
	}

	resp := new(freezonex_openiiot_api.StopApplicationResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// RestartApplication will restart application, if component_type is empty, then it will restart all component of this application
/*func (a *ApplicationService) RestartApplication(ctx context.Context, req *freezonex_openiiot_api.RestartApplicationRequest, c *app.RequestContext) (*freezonex_openiiot_api.RestartApplicationResponse, error) {

	_, tenantName, err := a.tenant.CheckTenant(ctx, req.TenantId, req.TenantName)
	if err != nil {
		logs.Error(ctx, "event=RestartApplication error=%v", err.Error())
		return nil, err
	}
	ctx = context.WithValue(ctx, "tenantName", tenantName)

	// check if application exist
	exist, err := a.ExistApplication(ctx, tenantName, req.ApplicationName)
	if err != nil {
		logs.Error(ctx, "event=Restartpplication error=%v", err.Error())
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("application name does not exist: %s", req.ApplicationName)
	}

	k8sUns := k8s.K8sUns{TenantName: tenantName, DeploymentCategory: "app", ComponentName: req.ApplicationName, ComponentType: req.ComponentType}

	deployments, err := a.k8s.GetDeploymentsByFuzzyName(ctx, k8sUns)
	if err != nil {
		logs.Error(ctx, "event=RestartApplication error=%v", err.Error())
		return nil, err
	}

	for _, deployment := range deployments {

		re := regexp.MustCompile(`^openiiot-app-([a-z0-9-]+)-(frontend|backend|db)$`)
		match := re.FindStringSubmatch(deployment.Name)
		if match == nil {
			return nil, fmt.Errorf("deployment name does not match the expected pattern")
		}
		k8sUns.ComponentType = match[2]

		a.k8s.RestartDeployment(ctx, k8sUns)
	}

	resp := new(freezonex_openiiot_api.RestartApplicationResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}*/

// DeleteApplication will delete application and related image, if component_type is empty, then it will delete all component of this application
func (a *ApplicationService) DeleteApplication(ctx context.Context, req *freezonex_openiiot_api.DeleteApplicationRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteApplicationResponse, error) {

	_, tenantName, err := a.tenant.CheckTenant(ctx, req.TenantId, req.TenantName)
	if err != nil {
		logs.Error(ctx, "event=DeleteApplication error=%v", err.Error())
		return nil, err
	}
	ctx = context.WithValue(ctx, "tenantName", tenantName)

	// check if application exist
	exist, err := a.ExistApplication(ctx, tenantName, req.ApplicationName)
	if err != nil {
		logs.Error(ctx, "event=DeleteApplication error=%v", err.Error())
		return nil, err
	}
	if !exist {
		return nil, fmt.Errorf("application name does not exist: %s", req.ApplicationName)
	}

	k8sUns := k8s.K8sUns{TenantName: tenantName, DeploymentCategory: "app", ComponentName: req.ApplicationName, ComponentType: req.ComponentType}

	deployments, err := a.k8s.GetDeploymentsByFuzzyName(ctx, k8sUns)
	if err != nil {
		logs.Error(ctx, "event=DeleteApplication error=%v", err.Error())
		return nil, err
	}

	for _, deployment := range deployments {
		/*if !strings.HasPrefix(deployment.Name, "openiiot-app-"+req.ApplicationName) {
			continue
		}*/
		re := regexp.MustCompile(`^openiiot-app-([a-z0-9-]+)-(frontend|backend|db)$`)
		match := re.FindStringSubmatch(deployment.Name)
		if match == nil {
			return nil, fmt.Errorf("deployment name does not match the expected pattern")
		}
		k8sUns.ComponentType = match[2]

		switch k8sUns.ComponentType {
		case "frontend":
			err = a.k8s.DeleteApplicationFrontend(ctx, k8sUns)
		case "backend":
			err = a.k8s.DeleteApplicationBackend(ctx, k8sUns)
		case "db":
			err = a.k8s.DeleteApplicationDb(ctx, k8sUns)
		}

		if err != nil {
			logs.Error(ctx, "event=DeleteApplication error=%v", err.Error())
			return nil, err
		}

		//delete docker image, ignore error
		err = a.DeleteDockerImage(a.k8s.GetAppImageFullName(k8sUns))
		if err != nil {
			logs.Error(ctx, "event=DeleteApplication error=%v", err.Error())
		}

		err = a.DeleteDockerImage(a.k8s.GetAppImageName(k8sUns))
		if err != nil {
			logs.Error(ctx, "event=DeleteApplication error=%v", err.Error())
		}

	}

	resp := new(freezonex_openiiot_api.DeleteApplicationResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}
