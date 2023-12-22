package main

import (
	//"context"
	//"os"

	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/dal/mysql"
	"freezonex/openiiot/biz/handler"
	"freezonex/openiiot/biz/middleware"
	iiotpb "freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/grafana"
	"freezonex/openiiot/biz/service/supos"
	"freezonex/openiiot/biz/service/tdengine"
	"freezonex/openiiot/biz/service/tenant"

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
	}

	tenantGroup := r.Group("/tenant", middleware.Access())
	{
		tenantHandler := handler.NewTenantHandler(tenant.NewTenantService(db))
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
	}

	userGroup := r.Group("/user", middleware.Access())
	{
		userHandler := handler.NewUserHandler(supos.NewSuposService(db, &c.SuposConfig))
		userGroup.GET(
			"/supos/list",
			middleware.Response(
				"/user/supos/list",
				userHandler.GetSuposUser,
				&iiotpb.GetSuposUserRequest{}))
		userGroup.GET(
			"/current",
			middleware.Response(
				"/user/current",
				userHandler.GetCurrentUser,
				&iiotpb.GetCurrentUserRequest{}))
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
			grafanaDashboardGroup.POST("/createdb", middleware.Response(
				"/grafana/dashboard/createdb",
				grafanaHandler.CreateDashBoard,
				&iiotpb.GrafanaCreateDashboardRequest{}))
			grafanaDashboardGroup.POST("/savebyuid", middleware.Response(
				"/grafana/dashboard/savebyuid",
				grafanaHandler.SaveDashboardByUid,
				&iiotpb.GrafanaSaveDashboardByUidRequest{}))
		}
		
		grafanaDataSourceGroup := grafanaGroup.Group("/datasource")
		{
			grafanaDataSourceGroup.GET("/getall", middleware.Response(
					"/grafana/datasource/getall",
					grafanaHandler.GetDatasources,
					&iiotpb.GrafanaDataSourcesRequest{}))
			grafanaDataSourceGroup.POST("/createds", middleware.Response(
					"/grafana/datasource/createds",
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

}
