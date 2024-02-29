package flow

import (
	"context"
	"fmt"
	"freezonex/openiiot/biz/dal/model_openiiot"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	application "freezonex/openiiot/biz/service/app"
	"freezonex/openiiot/biz/service/core"
	"freezonex/openiiot/biz/service/edge"
	"freezonex/openiiot/biz/service/emqx"
	"freezonex/openiiot/biz/service/grafana"
	"freezonex/openiiot/biz/service/tdengine"
	"freezonex/openiiot/biz/service/utils/common"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/json"
	"strings"
)

func (a *FlowService) AddFlow(ctx context.Context, req *freezonex_openiiot_api.AddFlowRequest, c *app.RequestContext) (*freezonex_openiiot_api.AddFlowResponse, error) {
	flowID, err := a.AddFlowDB(ctx, req.Name, req.Description, common.StringToInt64(req.TenantId), req.FlowType, "")
	if err != nil {
		logs.Error(ctx, "event=AddFlow error=%v", err.Error())
		return nil, err
	}

	if _, err := a.AddFlowEdge(ctx, flowID, common.StringToInt64Array(req.EdgeIds), "", "", "", ""); err != nil {
		return nil, err
	}
	if _, err := a.AddFlowCore(ctx, flowID, common.StringToInt64Array(req.CoreIds), "", ""); err != nil {
		return nil, err
	}
	if _, err := a.AddFlowApp(ctx, flowID, common.StringToInt64Array(req.AppIds), "", "", ""); err != nil {
		return nil, err
	}

	resp := new(freezonex_openiiot_api.AddFlowResponse)
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Id = common.Int64ToString(flowID)

	return resp, nil
}

// GetFlow will get flow record in condition
func (a *FlowService) GetFlow(ctx context.Context, req *freezonex_openiiot_api.GetFlowRequest, c *app.RequestContext) (*freezonex_openiiot_api.GetFlowResponse, error) {
	// Fetch flows from the database
	flows, err := a.GetFlowDB(ctx, common.StringToInt64(req.Id), req.Name, common.StringToInt64(req.TenantId), req.LastModifiedBy, req.FlowType)
	if err != nil {
		logs.Error(ctx, "event=GetFlow error=%v", err.Error())
		return nil, err
	}
	resp := new(freezonex_openiiot_api.GetFlowResponse)
	// Prepare the response data
	data := make([]*freezonex_openiiot_api.Flow, 0)
	for _, v := range flows {
		// For each flow, fetch EdgeIds, CoreIds, and AppIds from their respective tables
		edgeIds, err := a.GetFlowEdgeIDs(ctx, v.ID)
		if err != nil {
			return nil, err // Handle error appropriately
		}
		coreIds, err := a.GetFlowCoreIDs(ctx, v.ID)
		if err != nil {
			return nil, err // Handle error appropriately
		}
		appIds, err := a.GetFlowAppIDs(ctx, v.ID)
		if err != nil {
			return nil, err // Handle error appropriately
		}

		data = append(data, &freezonex_openiiot_api.Flow{
			Id:             common.Int64ToString(v.ID),
			Name:           v.Name,
			Description:    *v.Description,
			TenantId:       common.Int64ToString(v.TenantID),
			LastModifiedBy: *v.LastModifiedBy,
			FlowType:       *v.FlowType,
			EdgeIds:        common.Int64ToStringArray(edgeIds),
			CoreIds:        common.Int64ToStringArray(coreIds),
			AppIds:         common.Int64ToStringArray(appIds),
			CreateTime:     common.GetTimeStringFromTime(&v.CreateTime), // Format time as needed
			UpdateTime:     common.GetTimeStringFromTime(&v.UpdateTime),
		})
	}

	// Prepare and return the final response
	resp.BaseResp = middleware.SuccessResponseOK
	resp.Data = data

	return resp, nil
}

// UpdateFlow will update flow record
func (a *FlowService) UpdateFlow(ctx context.Context, req *freezonex_openiiot_api.UpdateFlowRequest, c *app.RequestContext) (*freezonex_openiiot_api.UpdateFlowResponse, error) {
	err := a.UpdateFlowDB(ctx, common.StringToInt64(req.Id), req.Name, req.Description, common.StringToInt64(req.TenantId), req.FlowType)

	if err != nil {
		logs.Error(ctx, "event=UpdateFlow error=%v", err.Error())
		return nil, err
	}
	// Update the flow_edge table
	err = a.UpdateFlowEdgeDB(ctx, common.StringToInt64(req.Id), req.EdgeIds)
	if err != nil {
		logs.Error(ctx, "event=UpdateFlowedge error=%v", err.Error())
		return nil, err
	}

	// Update the flow_core table
	err = a.UpdateFlowCoreDB(ctx, common.StringToInt64(req.Id), req.CoreIds)
	if err != nil {
		logs.Error(ctx, "event=UpdateFlowcore error=%v", err.Error())
		return nil, err
	}

	// Update the flow_app table
	err = a.UpdateFlowAppDB(ctx, common.StringToInt64(req.Id), req.AppIds)
	if err != nil {
		logs.Error(ctx, "event=UpdateFlowapp error=%v", err.Error())
		return nil, err
	}

	resp := new(freezonex_openiiot_api.UpdateFlowResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *FlowService) DeleteFlow(ctx context.Context, req *freezonex_openiiot_api.DeleteFlowRequest, c *app.RequestContext) (*freezonex_openiiot_api.DeleteFlowResponse, error) {
	// Delete related records in the flow_edge table
	if err := a.DeleteFlowEdgeRecords(ctx, common.StringToInt64(req.Id)); err != nil {
		logs.Error(ctx, "event=DeleteFlow error=%v", err.Error())
		return nil, err
	}

	// Delete related records in the flow_core table
	if err := a.DeleteFlowCoreRecords(ctx, common.StringToInt64(req.Id)); err != nil {
		logs.Error(ctx, "event=DeleteFlow error=%v", err.Error())
		return nil, err
	}

	// Delete related records in the flow_app table
	if err := a.DeleteFlowAppRecords(ctx, common.StringToInt64(req.Id)); err != nil {
		logs.Error(ctx, "event=DeleteFlow error=%v", err.Error())
		return nil, err
	}

	// Finally, delete the main flow record
	if err := a.DeleteFlowDB(ctx, common.StringToInt64(req.Id)); err != nil {
		logs.Error(ctx, "event=DeleteFlow error=%v", err.Error())
		return nil, err
	}

	// Prepare and return the response
	resp := new(freezonex_openiiot_api.DeleteFlowResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *FlowService) LoadDemoFlow(ctx context.Context, req *freezonex_openiiot_api.LoadDemoFlowRequest, c *app.RequestContext) (*freezonex_openiiot_api.LoadDemoFlowResponse, error) {

	err := a.LoadDemoFlowCore(ctx, req.DemoFlowId, req.TenantId)
	if err != nil {
		return nil, err
	}

	err = a.LoadDemoFlowApp(ctx, req.DemoFlowId, req.TenantId)
	if err != nil {
		return nil, err
	}
	err = a.LoadDemoFlowEdge(ctx, req.DemoFlowId, req.TenantId)
	if err != nil {
		return nil, err
	}

	resp := new(freezonex_openiiot_api.LoadDemoFlowResponse)
	resp.BaseResp = middleware.SuccessResponseOK

	return resp, nil
}

func (a *FlowService) LoadDemoFlowCore(ctx context.Context, DemoFlowId string, tenantid string) error {
	coreIds, err := a.GetFlowCoreIDs(ctx, common.StringToInt64(DemoFlowId))
	if err != nil {
		return err // Handle error appropriately
	}
	coreService := core.DefaultCoreService()
	if err != nil {
		return err // Handle error appropriately
	}
	var cores []*model_openiiot.Core

	// Iterate through coreIds to find a core of type "tdengine"
	for _, coreId := range coreIds {
		tempCores, err := coreService.GetCoreDB(ctx, coreId, "", 0, "", "tdengine")
		if err != nil {
			return err // Handle error appropriately, or continue to try the next ID
		}
		if len(tempCores) > 0 {
			cores = tempCores
			break // Found a core of type "tdengine", exit loop
		}
	}
	//外部测试使用端口
	if len(cores) == 0 {
		return fmt.Errorf("no cores of type 'tdengine' found") // or handle this case as appropriate
	}
	username := safeDeref(cores[0].Username)
	password := safeDeref(cores[0].Password)
	url := cores[0].URL
	//正式使用service+namespace
	//tenantService := tenant.DefaultTenantService()
	//tenant,err := tenantService.GetTenantDB(ctx,common.StringToInt64(tenantid),"")
	//url1 := fmt.Sprintf("tdengnine-openiiot-%s.openiiot-%ssvc.cluster.local:6041", tenant[0].Name, tenant[0].Name)

	tdengineService := tdengine.DefaultTdengineService()
	dsn := fmt.Sprintf("%s:%s@http(%s)/", username, password, url)
	//dsn := fmt.Sprintf("%s:%s@http(%s)/", "root", "taosdata", "47.236.10.165:30120")
	//tdengine-openiiot-shimu.openiiot-shimu.svc.cluster.local:6041 in inner pod

	scripts, _, _, err := a.GetFlowCoreScripts(ctx, cores[0].ID)
	if err != nil {
		return err
	}

	// Split the scripts into individual SQL commands
	sqlCommands := splitSQLCommands(scripts)

	for _, sqlCommand := range sqlCommands {
		if sqlCommand != "" {
			_, err = tdengineService.T.Exec(dsn, sqlCommand)
			logs.Info("sqlCommand", sqlCommand)
			if err != nil {
				// Handle execution error (you might want to log it or stop execution)
				return err
			}
		}
	}
	return nil
}

// splitSQLCommands splits a long SQL script into individual commands based on semicolons and line breaks.
func splitSQLCommands(scripts string) []string {
	// Normalize line endings to a single format (\n) to handle different OS line endings
	normalizedScripts := strings.ReplaceAll(scripts, "\r\n", "\n")
	normalizedScripts = strings.ReplaceAll(normalizedScripts, "\r", "\n")

	// Split the script by ";", allowing for potential trailing whitespace or new lines
	rawCommands := strings.Split(normalizedScripts, ";")

	var commands []string
	for _, command := range rawCommands {
		trimmedCommand := strings.TrimSpace(command)
		if trimmedCommand != "" {
			commands = append(commands, trimmedCommand+";")
		}
	}

	return commands
}

func safeDeref(ptr *string) string {
	if ptr != nil {
		return *ptr
	}
	return "" // or return a default value or handle this case as needed
}

func (a *FlowService) LoadDemoFlowApp(ctx context.Context, DemoFlowId string, tenantid string) error {

	appIds, err := a.GetFlowAppIDs(ctx, common.StringToInt64(DemoFlowId))
	if err != nil {
		return err // Handle error appropriately
	}
	appService := application.DefaultAppService()
	if err != nil {
		return err // Handle error appropriately
	}
	var apps []*model_openiiot.App

	// Iterate through app Ids to find a core of type "grafana"
	for _, appId := range appIds {
		tempApps, err := appService.GetAppDB(ctx, appId, "", 0, "", "grafana")
		if err != nil {
			return err // Handle error appropriately, or continue to try the next ID
		}
		if len(tempApps) > 0 {
			apps = tempApps
			break // Found a core of type "grafana", exit loop
		}
	}

	if len(apps) == 0 {
		return fmt.Errorf("no apps of type 'grafana' found") // or handle this case as appropriate
	}
	username := safeDeref(apps[0].Username)
	password := safeDeref(apps[0].Password)
	url := apps[0].URL
	scripts, scripts2, scripts3, err := a.GetFlowAppScripts(ctx, apps[0].ID)

	// Proceed with 'cores' which now contains cores of type "grafana"
	grafanaService := grafana.DefaultGrafanaService()

	//grafana-openiiot-shimu.openiiot-shimu.svc.cluster.local:3000 in inner pod

	// Construct the DSN part of the request
	dsn1 := &freezonex_openiiot_api.GrafanaDSN{
		Username: username,
		Password: password,
		Host:     url,
	}

	var dataSource *freezonex_openiiot_api.DataSource
	err = json.Unmarshal([]byte(scripts), &dataSource)
	if err != nil {
		return err // Handle JSON parsing error
	}

	req := &freezonex_openiiot_api.GrafanaCreateDataSourceRequest{
		DataSource: dataSource,
		Dsn:        dsn1,
	}

	if err != nil {
		return err
	}
	_, err = grafanaService.NewDataSource(ctx, req, nil)

	logs.Info("sqlCommand", scripts)
	if err != nil {
		// Handle execution error (you might want to log it or stop execution)
		return err
	}
	var dataSource2 *freezonex_openiiot_api.DataSource
	err = json.Unmarshal([]byte(scripts2), &dataSource2)
	if err != nil {
		return err // Handle JSON parsing error
	}

	req2 := &freezonex_openiiot_api.GrafanaCreateDataSourceRequest{
		DataSource: dataSource2,
		Dsn:        dsn1,
	}

	if err != nil {
		return err
	}
	_, err = grafanaService.NewDataSource(ctx, req2, nil)

	logs.Info("sqlCommand", scripts2)
	if err != nil {
		// Handle execution error (you might want to log it or stop execution)
		return err
	}

	var dashboard *freezonex_openiiot_api.Dashboard
	err = json.Unmarshal([]byte(scripts3), &dashboard)
	if err != nil {
		return err // Handle JSON parsing error
	}
	req3 := &freezonex_openiiot_api.GrafanaCreateDashboardRequest{
		Dashboard: dashboard,
		Dsn:       dsn1,
	}
	logs.Info("sqlCommand", scripts3)
	_, err = grafanaService.CreateDashBoard(ctx, req3, nil)
	if err != nil {
		return err // Handle JSON parsing error
	}
	return nil
}

func (a *FlowService) LoadDemoFlowEdge(ctx context.Context, DemoFlowId string, tenantid string) error {
	edgeIds, err := a.GetFlowEdgeIDs(ctx, common.StringToInt64(DemoFlowId))
	if err != nil {
		return err // Handle error appropriately
	}
	edgeService := edge.DefaultEdgeService()
	if err != nil {
		return err // Handle error appropriately
	}
	var edges []*model_openiiot.Edge

	for _, edgeId := range edgeIds {
		tempEdges, err := edgeService.GetEdgeDB(ctx, edgeId, "", 0, "", "emqx")
		if err != nil {
			return err // Handle error appropriately, or continue to try the next ID
		}
		if len(tempEdges) > 0 {
			edges = tempEdges
			break // Found a edge of type "emqx", exit loop
		}
	}
	//外部测试使用端口
	if len(edges) == 0 {
		return fmt.Errorf("no cores of type 'emqx' found") // or handle this case as appropriate
	}
	//username := safeDeref(edges[0].Username)
	//password := safeDeref(edges[0].Password)

	username := "85029670da5add00"
	password := "s6s9A9AtjAuQPIvDqDOkStuSUMJs9BeCgO0H0eIQGpD2uI"
	//username,password  need to replace

	url := edges[0].URL
	emqxService := emqx.DefaultEmqxService()
	//emqx-openiiot-shimu.openiiot-shimu.svc.cluster.local:18083 in inner pod

	scripts, scripts2, scripts3, scripts4, err := a.GetFlowEdgeScripts(ctx, edges[0].ID)
	if err != nil {
		return err
	}
	//grafana-openiiot-shimu.openiiot-shimu.svc.cluster.local:3000 in inner pod
	// Construct the DSN part of the request
	dsn1 := &freezonex_openiiot_api.EmqxDSN{
		Username: username,
		Password: password,
		Host:     url,
	}

	// Now create the GrafanaCreateDataSourceRequest

	var bridge *freezonex_openiiot_api.BridgeConfig
	err = json.Unmarshal([]byte(scripts), &bridge)
	if err != nil {
		return err // Handle JSON parsing error
	}

	req := &freezonex_openiiot_api.EmqxCreateBridgeRequest{
		Bridge: bridge,
		Dsn:    dsn1,
	}

	if err != nil {
		return err
	}
	_, err = emqxService.CreateBridge(ctx, req, nil)

	logs.Info("sqlCommand", scripts)
	if err != nil {
		return err
	}

	var bridge2 *freezonex_openiiot_api.BridgeConfig
	err = json.Unmarshal([]byte(scripts2), &bridge2)
	if err != nil {
		return err // Handle JSON parsing error
	}

	req2 := &freezonex_openiiot_api.EmqxCreateBridgeRequest{
		Bridge: bridge2,
		Dsn:    dsn1,
	}

	if err != nil {
		return err
	}
	_, err = emqxService.CreateBridge(ctx, req2, nil)

	logs.Info("sqlCommand", scripts2)
	if err != nil {
		return err
	}

	var rule *freezonex_openiiot_api.Rule
	err = json.Unmarshal([]byte(scripts3), &rule)
	if err != nil {
		return err // Handle JSON parsing error
	}

	req3 := &freezonex_openiiot_api.EmqxCreateRuleRequest{
		// Assuming the request requires a DataSource field populated from the struct we just unmarshaled
		Rule: rule,
		Dsn:  dsn1,
		// Populate othe,r required fields for the request here, such as Dsn if it's needed and different from the DataSource
	}
	logs.Info("sqlCommand", scripts3)
	_, err = emqxService.CreateRule(ctx, req3, nil)
	if err != nil {
		return err // Handle JSON parsing error
	}

	var rule2 *freezonex_openiiot_api.Rule
	err = json.Unmarshal([]byte(scripts4), &rule2)
	if err != nil {
		return err // Handle JSON parsing error
	}

	req4 := &freezonex_openiiot_api.EmqxCreateRuleRequest{
		Rule: rule2,
		Dsn:  dsn1,
	}
	logs.Info("sqlCommand", scripts4)
	_, err = emqxService.CreateRule(ctx, req4, nil)
	if err != nil {
		return err // Handle JSON parsing error
	}
	return nil
}
