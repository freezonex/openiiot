package tenant

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	application "freezonex/openiiot/biz/service/app"
	"freezonex/openiiot/biz/service/core"
	"freezonex/openiiot/biz/service/edge"
	"freezonex/openiiot/biz/service/flow"
	"freezonex/openiiot/biz/service/k8s"
	"freezonex/openiiot/biz/service/utils/common"
)

func (a *TenantService) AddTenant(ctx context.Context, req *freezonex_openiiot_api.AddTenantRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddTenantResponse, error) {
	_, tenantName, err := a.CheckTenant(ctx, "", req.Name)
	if err == nil {
		return nil, fmt.Errorf("tenant already exist in database: %s", tenantName)
	}

	tenantID, err := a.AddTenantDB(ctx, req.Name, req.Description)
	if err != nil {
		logs.Error(ctx, "event=AddTenant error=%v", err.Error())
		return nil, err
	}

	//tenantID, tenantName, _ := a.CheckTenant(ctx, "", req.Name)

	tenantName = req.Name
	ctx = context.WithValue(ctx, "tenantName", tenantName)

	// err = a.AddDefaultFlow(ctx, tenantID, req.Name)

	/*if tenantName == "dt" {
		if err := a.k8s.CreateNamespace(ctx, k8s.K8sUns{TenantName: tenantName}); err != nil {
			return nil, err
		}
		if err := a.K8sBasicInstall(ctx); err != nil {
			return nil, err
		}
	}*/

	if err := a.K8sTenantAdd(ctx, tenantName); err != nil {
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddTenantResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = common.Int64ToString(tenantID)

	return resp, nil
}

func (a *TenantService) K8sBasicInstall(ctx context.Context) error {

	if err := a.k8s.CreateNamespace(ctx, k8s.K8sUns{}); err != nil {
		return err
	}

	k8sUns := k8s.K8sUns{ComponentName: "mysql"}
	if err := a.k8s.CreateMysqlComponent(ctx, k8sUns); err != nil {
		return err
	}

	k8sUns = k8s.K8sUns{ComponentName: "server"}
	if err := a.k8s.CreateMysqlComponent(ctx, k8sUns); err != nil {
		return err
	}

	k8sUns = k8s.K8sUns{ComponentName: "consolemanager"}
	if err := a.k8s.CreateConsolemanagerComponent(ctx, k8sUns); err != nil {
		return err
	}

	return nil
}

func (a *TenantService) K8sTenantAdd(ctx context.Context, tenantName string) error {

	// check if tenant already exist
	k8sUns := k8s.K8sUns{TenantName: tenantName}
	namespaceName := a.k8s.GetNamespaceName(k8sUns)
	exist, err := a.k8s.NamespaceExists(ctx, namespaceName)
	if err != nil {
		return fmt.Errorf("failed to check namespace existence: %w", err)
	}
	if exist {
		return fmt.Errorf("tenant namespace already exist: %s", namespaceName)
	}

	if err := a.k8s.CreateNamespace(ctx, k8sUns); err != nil {
		return err
	}

	k8sUns = k8s.K8sUns{TenantName: tenantName, ComponentName: "nodered", Number: "1"}
	if err := a.k8s.CreateNoderedComponent(ctx, k8sUns); err != nil {
		return err
	}

	k8sUns = k8s.K8sUns{TenantName: tenantName, ComponentName: "grafana", Number: "1"}
	if err := a.k8s.CreateGrafanaComponent(ctx, k8sUns); err != nil {
		return err
	}

	k8sUns = k8s.K8sUns{TenantName: tenantName, ComponentName: "emqx"}
	if err := a.k8s.CreateEmqxComponent(ctx, k8sUns); err != nil {
		return err
	}

	k8sUns = k8s.K8sUns{TenantName: tenantName, ComponentName: "tdengine"}
	if err := a.k8s.CreateTdengineComponent(ctx, k8sUns); err != nil {
		return err
	}

	return nil
}

func (a *TenantService) K8sTenantDelete(ctx context.Context, tenantName string) error {

	// check if tenant already exist
	k8sUns := k8s.K8sUns{TenantName: tenantName}
	namespaceName := a.k8s.GetNamespaceName(k8sUns)
	exist, err := a.k8s.NamespaceExists(ctx, namespaceName)
	if err != nil {
		return fmt.Errorf("failed to check namespace existence: %w", err)
	}
	if !exist {
		logs.Error(ctx, "event=DeleteTenant namespace not exists%s", namespaceName)
		return nil
	}

	if err := a.k8s.DeleteNamespace(ctx, k8sUns); err != nil {
		return err
	}

	// pv not belong to any namespace, need delete separately
	if err := a.k8s.DeleteTenantPersistentVolume(ctx, tenantName); err != nil {
		return err
	}

	runtime_idc_name := os.Getenv("RUNTIME_IDC_NAME")
	if runtime_idc_name != "ci" && runtime_idc_name != "local" {
		if err := a.k8s.DeleteTenantPersistentVolumePath(ctx, tenantName); err != nil {
			return err
		}
	}

	return err
}

// GetTenant will get tenant record in condition
func (a *TenantService) GetTenant(ctx context.Context, req *freezonex_openiiot_api.GetTenantRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetTenantResponse, error) {
	tenants, err := a.GetTenantDB(ctx, common.StringToInt64(req.Id), req.Name)

	if err != nil {
		logs.Error(ctx, "event=GetTenant error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetTenantResponse)
	data := make([]*freezonex_openiiot_api.Tenant, 0)
	for _, v := range tenants {
		data = append(data, &freezonex_openiiot_api.Tenant{
			Id:          common.Int64ToString(v.ID),
			Name:        v.Name,
			Description: *v.Description,
			IsDefault:   common.ProtectNullBoolPointer(v.IsDefault),
			CreateTime:  common.GetTimeStringFromTime(&v.CreateTime), // Format time as needed
			UpdateTime:  common.GetTimeStringFromTime(&v.UpdateTime),
		})
	}
	resp.Data = data
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// UpdateTenant will update tenant record
func (a *TenantService) UpdateTenant(ctx context.Context, req *freezonex_openiiot_api.UpdateTenantRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateTenantResponse, error) {
	err := a.UpdateTenantDB(ctx, common.StringToInt64(req.Id), req.Name, req.Description)
	if err != nil {
		logs.Error(ctx, "event=UpdateTenant error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateTenantResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

// DeleteTenant will delete tenant namespace from k8s also
func (a *TenantService) DeleteTenant(ctx context.Context, req *freezonex_openiiot_api.DeleteTenantRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteTenantResponse, error) {
	//Delete tenant also should delete tenant user, edge pool, core pool, application pool, flow
	/*err := a.DeleteTenantUserDB(ctx, req.Id)
	if err != nil {
		logs.Error(ctx, "event=DeleteTenant user error=%v", err.Error())
		return nil, err
	}*/
	//delete all the users and flows?

	_, tenantName, err := a.CheckTenant(ctx, req.Id, req.Name)
	if err != nil {
		logs.Error(ctx, "event=CheckTenant error=%v", err.Error())
		return nil, err
	}
	if tenantName == "" {
		return nil, fmt.Errorf("tenant not exist in database: %s", tenantName)
	}

	runtime_idc_name := os.Getenv("RUNTIME_IDC_NAME")
	if runtime_idc_name != "ci" && runtime_idc_name != "local" {
		if tenantName == "dt" {
			return nil, errors.New("Default tenant cannot be deleted")
		}
	}

	if err := a.K8sTenantDelete(ctx, req.Name); err != nil {
		return nil, err
	}

	if err := a.DeleteTenantDB(ctx, tenantName); err != nil {
		return nil, err
	}

	resp := new(freezonex_openiiot_api.DeleteTenantResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *TenantService) GetAllTenantName(ctx context.Context, req *freezonex_openiiot_api.GetAllTenantNameRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetAllTenantNameResponse, error) {
	tenants, err := a.GetTenantDB(ctx, 0, "")

	if err != nil {
		logs.Error(ctx, "event=GetTenant error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.GetAllTenantNameResponse)
	tenantNames := make([]string, 0) // Change to slice of strings
	for _, v := range tenants {
		tenantNames = append(tenantNames, v.Name) // Append only the name
	}

	resp.TenantNames = tenantNames
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *TenantService) AddDefaultFlow(ctx context.Context, tenantid int64, name string) error {
	edgeService := edge.DefaultEdgeService()
	emqxname := "emqx-openiiot-" + name
	emqxurl := "http://emqx-openiiot-" + name + ".openiiot-" + name + ".svc.cluster.local:18083"
	edgeid, err := edgeService.AddEdgeDB(ctx, emqxname, "", tenantid, emqxurl, "admin", "public", "emqx")
	if err != nil {
		return err
	}
	var edgeids []int64
	edgeids = append(edgeids, edgeid)

	coreService := core.DefaultCoreService()
	tdenginename := "tdengine-openiiot-" + name
	tdengineurl := "http://tdengine-openiiot-" + name + ".openiiot-" + name + ".svc.cluster.local:6041"
	coreid, err := coreService.AddCoreDB(ctx, tdenginename, "", tenantid, tdengineurl, "root", "taosdata", "tdengine")
	if err != nil {
		return err
	}
	var coreids []int64
	coreids = append(coreids, coreid)

	appService := application.DefaultAppService()
	grafananame := "grafana-openiiot-" + name
	grafanaurl := "http://grafana-openiiot-" + name + ".openiiot-" + name + ".svc.cluster.local:3000"
	appid, err := appService.AddAppDB(ctx, grafananame, "", tenantid, grafanaurl, "admin", "admin", "grafana")
	if err != nil {
		return err
	}
	var appids []int64
	appids = append(appids, appid)

	flowService := flow.DefaultFlowService()
	flowname := "flow-openiiot-" + name
	flowid, err := flowService.AddFlowDB(ctx, flowname, "demoflow", tenantid, "demoflow", "system")

	edgescripts1json := emqxreplace1(name)
	edgescripts1, err := json.MarshalIndent(edgescripts1json, "", "  ")
	if err != nil {
		return err
	}

	edgescripts2json := emqxreplace2(name)
	edgescripts2, err := json.MarshalIndent(edgescripts2json, "", "  ")
	if err != nil {
		return err
	}
	edgescripts3 := staticText1
	edgescripts4 := staticText2

	_, err = flowService.AddFlowEdge(ctx, flowid, edgeids, string(edgescripts1), string(edgescripts2), edgescripts3, edgescripts4)
	if err != nil {
		return err
	}

	corescripts1 := staticText3
	_, err = flowService.AddFlowCore(ctx, flowid, coreids, corescripts1, "")
	if err != nil {
		return err
	}

	appscripts1json := grafanareplace1(name)
	appscripts1, err := json.MarshalIndent(appscripts1json, "", "  ")
	appscripts2json := grafanareplace2(name)
	appscripts2, err := json.MarshalIndent(appscripts2json, "", "  ")
	appscripts3 := staticText4

	_, err = flowService.AddFlowApp(ctx, flowid, appids, string(appscripts1), string(appscripts2), appscripts3)
	if err != nil {
		return err
	}

	return nil
}
