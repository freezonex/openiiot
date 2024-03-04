package main

import (
	//"context"
	//"os"

	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/mysql"
	"freezonex/openiiot/biz/handler"
	"freezonex/openiiot/biz/middleware"
	iiotpb "freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/app"
	"freezonex/openiiot/biz/service/core"
	"freezonex/openiiot/biz/service/edge"
	"freezonex/openiiot/biz/service/emqx"
	"freezonex/openiiot/biz/service/flow"
	"freezonex/openiiot/biz/service/grafana"
	"freezonex/openiiot/biz/service/supos"
	"freezonex/openiiot/biz/service/tdengine"
	"freezonex/openiiot/biz/service/tenant"
	"freezonex/openiiot/biz/service/user"
	"freezonex/openiiot/biz/service/wms_warehouse"

	"github.com/cloudwego/hertz/pkg/app/server"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

// customizeRegister register customize routers.
func customizeRegister(r *server.Hertz, c *config.Config) {
	db, err := mysql.Init(&c.DBConfig)
	if err != nil {
		logs.Errorf("cannot connect to mysql database, err: %v", err)
		panic(err)
	}

	authGroup := r.Group("/auth", middleware.Access())
	{
		authHandler := handler.NewAuthHandler(supos.NewSuposService(db, &c.SuposConfig))
		authGroup.GET("/authorize", authHandler.Authorize)
		authGroup.GET(
			"/accessToken",
			middleware.Response(
				"/auth/accessToken",
				authHandler.AccessToken,
				&iiotpb.GetAccessTokenRequest{}))
		authGroup.GET(
			"/logout",
			middleware.Response(
				"/auth/logout",
				authHandler.Logout,
				&iiotpb.LogoutRequest{}))
		authGroup.POST(
			"/login",
			middleware.Response(
				"/auth/login",
				authHandler.Login,
				&iiotpb.LoginRequest{}))
	}

	tenantGroup := r.Group("/tenant", middleware.Access())
	{
		tenantHandler := handler.NewTenantHandler(tenant.NewTenantService(db, &c.K8sConfig))
		tenantGroup.POST(
			"/add",
			middleware.Response(
				"/tenant/add",
				tenantHandler.AddTenant,
				&iiotpb.AddTenantRequest{}))
		tenantGroup.GET(
			"/get",
			middleware.Response(
				"/tenant/get",
				tenantHandler.GetTenant,
				&iiotpb.GetTenantRequest{}))
		tenantGroup.POST(
			"/update",
			middleware.Response(
				"/tenant/update",
				tenantHandler.UpdateTenant,
				&iiotpb.UpdateTenantRequest{}))
		tenantGroup.POST(
			"/delete",
			middleware.Response(
				"/tenant/delete",
				tenantHandler.DeleteTenant,
				&iiotpb.DeleteTenantRequest{}))
		tenantGroup.GET(
			"/allname",
			middleware.Response(
				"/tenant/allname",
				tenantHandler.GetAllTenantName,
				&iiotpb.GetAllTenantNameRequest{}))
	}

	userGroup := r.Group("/user", middleware.Access())
	{
		suposuserHandler := handler.NewSuposHandler(supos.NewSuposService(db, &c.SuposConfig))
		userHandler := handler.NewUserHandler(user.NewUserService(db))
		userGroup.GET(
			"/supos/list",
			middleware.Response(
				"/user/supos/list",
				suposuserHandler.GetSuposUser,
				&iiotpb.GetSuposUserRequest{}))
		userGroup.GET(
			"/current",
			middleware.Response(
				"/user/current",
				suposuserHandler.GetCurrentUser,
				&iiotpb.GetCurrentUserRequest{}))
		userGroup.POST(
			"/add",
			middleware.Response(
				"/user/add",
				userHandler.AddUser,
				&iiotpb.AddUserRequest{}))
		userGroup.GET(
			"/get",
			middleware.Response(
				"/user/get",
				userHandler.GetUser,
				&iiotpb.GetUserRequest{}))
		userGroup.POST(
			"/update",
			middleware.Response(
				"/user/update",
				userHandler.UpdateUser,
				&iiotpb.UpdateUserRequest{}))
		userGroup.POST(
			"/delete",
			middleware.Response(
				"/user/delete",
				userHandler.DeleteUser,
				&iiotpb.DeleteUserRequest{}))
		userGroup.POST(
			"/getuserbytoken",
			middleware.Response(
				"/user/getuserbytoken",
				userHandler.GetUserByToken,
				&iiotpb.GetUserByTokenRequest{}))
	}

	edgeGroup := r.Group("/edge", middleware.Access())
	{
		edgeHandler := handler.NewEdgeHandler(edge.NewEdgeService(db))
		edgeGroup.POST(
			"/add",
			middleware.Response(
				"/edge/add",
				edgeHandler.AddEdge,
				&iiotpb.AddEdgeRequest{}))
		edgeGroup.GET(
			"/get",
			middleware.Response(
				"/edge/get",
				edgeHandler.GetEdge,
				&iiotpb.GetEdgeRequest{}))
		edgeGroup.POST(
			"/update",
			middleware.Response(
				"/edge/update",
				edgeHandler.UpdateEdge,
				&iiotpb.UpdateEdgeRequest{}))
		edgeGroup.POST(
			"/delete",
			middleware.Response(
				"/edge/delete",
				edgeHandler.DeleteEdge,
				&iiotpb.DeleteEdgeRequest{}))
	}

	coreGroup := r.Group("/core", middleware.Access())
	{
		coreHandler := handler.NewCoreHandler(core.NewCoreService(db))
		coreGroup.POST(
			"/add",
			middleware.Response(
				"/core/add",
				coreHandler.AddCore,
				&iiotpb.AddCoreRequest{}))
		coreGroup.GET(
			"/get",
			middleware.Response(
				"/core/get",
				coreHandler.GetCore,
				&iiotpb.GetCoreRequest{}))
		coreGroup.POST(
			"/update",
			middleware.Response(
				"/core/update",
				coreHandler.UpdateCore,
				&iiotpb.UpdateCoreRequest{}))
		coreGroup.POST(
			"/delete",
			middleware.Response(
				"/core/delete",
				coreHandler.DeleteCore,
				&iiotpb.DeleteCoreRequest{}))
	}

	appGroup := r.Group("/app", middleware.Access())
	{
		appHandler := handler.NewAppHandler(app.NewAppService(db))
		appGroup.POST(
			"/add",
			middleware.Response(
				"/app/add",
				appHandler.AddApp,
				&iiotpb.AddAppRequest{}))
		appGroup.GET(
			"/get",
			middleware.Response(
				"/app/get",
				appHandler.GetApp,
				&iiotpb.GetAppRequest{}))
		appGroup.POST(
			"/update",
			middleware.Response(
				"/app/update",
				appHandler.UpdateApp,
				&iiotpb.UpdateAppRequest{}))
		appGroup.POST(
			"/delete",
			middleware.Response(
				"/app/delete",
				appHandler.DeleteApp,
				&iiotpb.DeleteAppRequest{}))
	}

	flowGroup := r.Group("/flow", middleware.Access())
	{
		flowHandler := handler.NewFlowHandler(flow.NewFlowService(db))
		flowGroup.POST(
			"/add",
			middleware.Response(
				"/flow/add",
				flowHandler.AddFlow,
				&iiotpb.AddFlowRequest{}))
		flowGroup.GET(
			"/get",
			middleware.Response(
				"/flow/get",
				flowHandler.GetFlow,
				&iiotpb.GetFlowRequest{}))
		flowGroup.POST(
			"/update",
			middleware.Response(
				"/flow/update",
				flowHandler.UpdateFlow,
				&iiotpb.UpdateFlowRequest{}))
		flowGroup.POST(
			"/delete",
			middleware.Response(
				"/flow/delete",
				flowHandler.DeleteFlow,
				&iiotpb.DeleteFlowRequest{}))
		flowGroup.POST(
			"/loaddemoflow",
			middleware.Response(
				"/flow/loaddemoflow",
				flowHandler.LoadDemoFlow,
				&iiotpb.LoadDemoFlowRequest{}))
	}

	emqxGroup := r.Group("/emqx", middleware.Access())
	{
		emqxHandler := handler.NewEmqxHandler(emqx.NewEmqxService(db))
		emqxGroup.GET(
			"/status",
			middleware.Response(
				"/emqx/status",
				emqxHandler.GetStatus,
				&iiotpb.EmqxGetStatusRequest{}))
		emqxBridgeGroup := emqxGroup.Group("/bridge")
		{
			emqxBridgeGroup.POST("/create",
				middleware.Response(
					"/emqx/bridge/create",
					emqxHandler.CreateBridge,
					&iiotpb.EmqxCreateBridgeRequest{}))
		}
		emqxRuleGroup := emqxGroup.Group("/rule")
		{
			emqxRuleGroup.POST("/create",
				middleware.Response(
					"/emqx/rule/create",
					emqxHandler.CreateRule,
					&iiotpb.EmqxCreateRuleRequest{}))
		}
	}

	grafanaGroup := r.Group("/grafana", middleware.Access())
	{
		grafanaHandler := handler.NewGrafanaHandler(grafana.NewGrafanaService(db, &c.GrafanaConfig))
		grafanaAuthGroup := grafanaGroup.Group("/auth")
		{
			grafanaAuthGroup.GET("/authorize", grafanaHandler.Authorize)
			grafanaAuthGroup.POST(
				"/accesstoken",
				middleware.Response(
					"/grafana/auth/accesstoken",
					grafanaHandler.GetAccessToken,
					&iiotpb.GrafanaAccessTokenRequest{}))
		}

		grafanaGroup.GET(
			"/user",
			middleware.Response(
				"/grafana/user",
				grafanaHandler.GetUser,
				&iiotpb.GrafanaUserRequest{}))
		grafanaDashboardGroup := grafanaGroup.Group("/dashboard")
		{
			grafanaDashboardGroup.POST("/create", middleware.Response(
				"/grafana/dashboard/create",
				grafanaHandler.CreateDashBoard,
				&iiotpb.GrafanaCreateDashboardRequest{}))
			grafanaDashboardGroup.POST("/save/uid", middleware.Response(
				"/grafana/dashboard/save/uid",
				grafanaHandler.SaveDashboardByUid,
				&iiotpb.GrafanaSaveDashboardByUidRequest{}))
		}

		grafanaDataSourceGroup := grafanaGroup.Group("/datasource")
		{
			grafanaDataSourceGroup.GET("/get", middleware.Response(
				"/grafana/datasource/get",
				grafanaHandler.GetDatasource,
				&iiotpb.GrafanaGetDataSourceRequest{}))
			grafanaDataSourceGroup.POST("/create", middleware.Response(
				"/grafana/datasource/create",
				grafanaHandler.CreateDatasource,
				&iiotpb.GrafanaCreateDataSourceRequest{}))
			grafanaDataSourceGroup.POST("/delete", middleware.Response(
				"/grafana/datasource/delete",
				grafanaHandler.DeleteDatasource,
				&iiotpb.GrafanaDeleteDataSourceRequest{}))
		}

	}

	tdengineGroup := r.Group("/tdengine", middleware.Access())
	{
		tdengineHandler := handler.NewTDEngineHandler(tdengine.NewTDEngineService(db, &c.TDEngineConfig))
		tdengineGroup.POST(
			"/testconnection",
			middleware.Response(
				"/tdengine/testconnection",
				tdengineHandler.TestConnection,
				&iiotpb.TDEngineTestConnectionRequest{}))
		tdengineGroup.POST(
			"/exec",
			middleware.Response(
				"/tdengine/exec",
				tdengineHandler.Exec,
				&iiotpb.TDEngineExecRequest{}))
		tdengineGroup.POST(
			"/batchexec",
			middleware.Response(
				"/tdengine/batchexec",
				tdengineHandler.BatchExec,
				&iiotpb.TDEngineBatchExecRequest{}))
		tdengineGroup.POST(
			"/query",
			middleware.Response(
				"/tdengine/query",
				tdengineHandler.Query,
				&iiotpb.TDEngineQueryRequest{}))
	}

	wmsGroup := r.Group("/wmswarehouse", middleware.Access())
	{
		wmsHandler := handler.NewWmsWarehouseHandler(wms_warehouse.NewWmsWarehouseService(db))
		wmsGroup.POST(
			"/add",
			middleware.Response(
				"/tenant/add",
				wmsHandler.AddWmsWarehouse,
				&iiotpb.TDEngineTestConnectionRequest{}))
		wmsGroup.POST(
			"/update",
			middleware.Response(
				"/tenant/update",
				wmsHandler.UpdateWmsWarehouse,
				&iiotpb.TDEngineExecRequest{}))
		wmsGroup.POST(
			"/delete",
			middleware.Response(
				"/tenant/delete",
				wmsHandler.DeleteWmsWarehouse,
				&iiotpb.TDEngineBatchExecRequest{}))
		wmsGroup.POST(
			"/get",
			middleware.Response(
				"/tenant/get",
				wmsHandler.GetWmsWarehouse,
				&iiotpb.TDEngineQueryRequest{}))
	}

}
