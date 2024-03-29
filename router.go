package main

import (
	"freezonex/openiiot/biz/service/wms_rfid_material"
	//"context"
	//"os"
	"github.com/cloudwego/hertz/pkg/app/server"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

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
	"freezonex/openiiot/biz/service/wms_inbound"
	"freezonex/openiiot/biz/service/wms_inbound_record"
	"freezonex/openiiot/biz/service/wms_material"
	"freezonex/openiiot/biz/service/wms_outbound"
	"freezonex/openiiot/biz/service/wms_outbound_record"
	"freezonex/openiiot/biz/service/wms_stocktaking"
	"freezonex/openiiot/biz/service/wms_stocktaking_record"
	"freezonex/openiiot/biz/service/wms_storage_location"
	"freezonex/openiiot/biz/service/wms_storagelocationmaterial"
	"freezonex/openiiot/biz/service/wms_warehouse"
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

	wmsWarehouseGroup := r.Group("/wms/warehouse", middleware.Access())
	{
		wmsHandler := handler.NewWmsWarehouseHandler(wms_warehouse.NewWmsWarehouseService(db))
		wmsWarehouseGroup.POST(
			"/add",
			middleware.Response(
				"/wms/warehouse/add",
				wmsHandler.AddWmsWarehouse,
				&iiotpb.AddWarehouseRequest{}))
		wmsWarehouseGroup.POST(
			"/update",
			middleware.Response(
				"/wms/warehouse/update",
				wmsHandler.UpdateWmsWarehouse,
				&iiotpb.UpdateWarehouseRequest{}))
		wmsWarehouseGroup.POST(
			"/delete",
			middleware.Response(
				"/wms/warehouse/delete",
				wmsHandler.DeleteWmsWarehouse,
				&iiotpb.DeleteWarehouseRequest{}))
		wmsWarehouseGroup.POST(
			"/get",
			middleware.Response(
				"/wms/warehouse/get",
				wmsHandler.GetWmsWarehouse,
				&iiotpb.GetWarehouseRequest{}))
	}

	wmsMaterialGroup := r.Group("/wms/material", middleware.Access())
	{
		wms_storagelocationmaterial.NewStorageLocationMaterial(db)
		wmsHandler := handler.NewWmsMaterialHandler(wms_material.NewWmsMaterialService(db))
		wmsMaterialGroup.POST(
			"/add",
			middleware.Response(
				"/wms/material/add",
				wmsHandler.AddWmsMaterial,
				&iiotpb.AddMaterialRequest{}))
		wmsMaterialGroup.POST(
			"/update",
			middleware.Response(
				"/wms/material/update",
				wmsHandler.UpdateWmsMaterial,
				&iiotpb.UpdateMaterialRequest{}))
		wmsMaterialGroup.POST(
			"/delete",
			middleware.Response(
				"/wms/material/delete",
				wmsHandler.DeleteWmsMaterial,
				&iiotpb.DeleteMaterialRequest{}))
		wmsMaterialGroup.POST(
			"/get",
			middleware.Response(
				"/wms/material/get",
				wmsHandler.GetWmsMaterial,
				&iiotpb.GetMaterialRequest{}))
	}
	wmsStorageLocationGroup := r.Group("/wms/storagelocation", middleware.Access())
	{
		wmsHandler := handler.NewWmsStorageLocationHandler(wms_storage_location.NewStorageLocationService(db))
		wmsStorageLocationGroup.POST(
			"/add",
			middleware.Response(
				"/wms/storagelocation/add",
				wmsHandler.AddWmsStorageLocation,
				&iiotpb.AddStorageLocationRequest{}))
		wmsStorageLocationGroup.POST(
			"/update",
			middleware.Response(
				"/wms/storagelocation/update",
				wmsHandler.UpdateWmsStorageLocation,
				&iiotpb.UpdateStorageLocationRequest{}))
		wmsStorageLocationGroup.POST(
			"/delete",
			middleware.Response(
				"/wms/storagelocation/delete",
				wmsHandler.DeleteWmsStorageLocation,
				&iiotpb.DeleteStorageLocationRequest{}))
		wmsStorageLocationGroup.POST(
			"/get",
			middleware.Response(
				"/wms/storagelocation/get",
				wmsHandler.GetWmsStorageLocation,
				&iiotpb.GetStorageLocationRequest{}))
	}

	wmsInboundGroup := r.Group("/wms/inbound", middleware.Access())
	{
		wmsHandler := handler.NewWmsInboundHandler(wms_inbound.NewWmsInboundService(db))
		wmsInboundGroup.POST(
			"/add",
			middleware.Response(
				"/wms/inbound/add",
				wmsHandler.AddWmsInbound,
				&iiotpb.AddInboundRequest{}))
		wmsInboundGroup.POST(
			"/update",
			middleware.Response(
				"/wms/inbound/update",
				wmsHandler.UpdateWmsInbound,
				&iiotpb.UpdateInboundRequest{}))
		wmsInboundGroup.POST(
			"/delete",
			middleware.Response(
				"/wms/inbound/delete",
				wmsHandler.DeleteWmsInbound,
				&iiotpb.DeleteInboundRequest{}))
		wmsInboundGroup.POST(
			"/get",
			middleware.Response(
				"/wms/inbound/get",
				wmsHandler.GetWmsInbound,
				&iiotpb.GetInboundRequest{}))
	}

	wmsOutboundGroup := r.Group("/wms/outbound", middleware.Access())
	{
		wmsHandler := handler.NewWmsOutboundHandler(wms_outbound.NewWmsOutboundService(db))
		wmsOutboundGroup.POST(
			"/add",
			middleware.Response(
				"/wms/outbound/add",
				wmsHandler.AddWmsOutbound,
				&iiotpb.AddOutboundRequest{}))
		wmsOutboundGroup.POST(
			"/update",
			middleware.Response(
				"/wms/outbound/update",
				wmsHandler.UpdateWmsOutbound,
				&iiotpb.UpdateOutboundRequest{}))
		wmsOutboundGroup.POST(
			"/delete",
			middleware.Response(
				"/wms/outbound/delete",
				wmsHandler.DeleteWmsOutbound,
				&iiotpb.DeleteOutboundRequest{}))
		wmsOutboundGroup.POST(
			"/get",
			middleware.Response(
				"/wms/outbound/get",
				wmsHandler.GetWmsOutbound,
				&iiotpb.GetOutboundRequest{}))
	}

	wmsRfidmaterialGroup := r.Group("/wms/rfidmaterial", middleware.Access())
	{
		wmsHandler := handler.NewWmsRfidMaterialHandler(wms_rfid_material.NewWmsRfidMaterialService(db))
		wmsRfidmaterialGroup.POST(
			"/add",
			middleware.Response(
				"/wms/rfidmaterial/add",
				wmsHandler.AddWmsRfidMaterial,
				&iiotpb.AddRfidMaterialRequest{}))
		wmsRfidmaterialGroup.POST(
			"/update",
			middleware.Response(
				"/wms/rfidmaterial/update",
				wmsHandler.UpdateWmsRfidMaterial,
				&iiotpb.UpdateRfidMaterialRequest{}))
		wmsRfidmaterialGroup.POST(
			"/delete",
			middleware.Response(
				"/wms/rfidmaterial/delete",
				wmsHandler.DeleteWmsRfidMaterial,
				&iiotpb.DeleteRfidMaterialRequest{}))
		wmsRfidmaterialGroup.POST(
			"/get",
			middleware.Response(
				"/wms/rfidmaterial/get",
				wmsHandler.GetWmsRfidMaterial,
				&iiotpb.GetRfidMaterialRequest{}))
	}

	wmsInboundRecordGroup := r.Group("/wms/inbound/detail", middleware.Access())
	{
		wmsHandler1 := handler.NewWmsInboundRecordHandler(wms_inbound_record.NewWmsInboundRecordService(db))

		wmsInboundRecordGroup.POST(
			"/get",
			middleware.Response(
				"/wms/inbound/detail/get",
				wmsHandler1.GetWmsInboundRecord,
				&iiotpb.GetInboundRecordRequest{}))
	}

	wmsOutboundRecordGroup := r.Group("/wms/outbound/detail", middleware.Access())
	{
		wmsHandler1 := handler.NewWmsOutboundRecordHandler(wms_outbound_record.NewWmsOutboundRecordService(db))

		wmsOutboundRecordGroup.POST(
			"/get",
			middleware.Response(
				"/wms/outbound/detail/get",
				wmsHandler1.GetWmsOutboundRecord,
				&iiotpb.GetOutboundRecordRequest{}))
	}
	wms_StocktakingGroup := r.Group("/wms/stocktaking", middleware.Access())
	{
		wmsHandler := handler.NewWmsStocktakingHandler(wms_stocktaking.NewStocktakingService(db))
		wms_StocktakingGroup.POST(
			"/add",
			middleware.Response(
				"/wms/stocktaking/add",
				wmsHandler.AddWmsStocktaking,
				&iiotpb.AddStocktakingRequest{}))
		wms_StocktakingGroup.POST(
			"/update",
			middleware.Response(
				"/wms/stocktaking/update",
				wmsHandler.UpdateWmsStocktaking,
				&iiotpb.UpdateStocktakingRequest{}))
		wms_StocktakingGroup.POST(
			"/delete",
			middleware.Response(
				"/wms/stocktaking/delete",
				wmsHandler.DeleteWmsStocktaking,
				&iiotpb.DeleteStocktakingRequest{}))
		wms_StocktakingGroup.POST(
			"/get",
			middleware.Response(
				"/wms/stocktaking/get",
				wmsHandler.GetWmsStocktaking,
				&iiotpb.GetStocktakingRequest{}))
	}

	wms_Stocktaking_recordGroup := r.Group("/wms/stocktaking/detail", middleware.Access())
	{
		wmsHandler := handler.NewWmsStocktakingRecordHandler(wms_stocktaking_record.NewStocktakingRecordService(db))
		wms_Stocktaking_recordGroup.POST(
			"/get",
			middleware.Response(
				"/wms/stocktaking/detail/get",
				wmsHandler.GetWmsStocktakingRecord,
				&iiotpb.GetStocktakingRecordRequest{}))
	}
}
