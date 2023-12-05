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

	"github.com/cloudwego/hertz/pkg/app/server"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

// customizeRegister register customize routers.
func customizeRegister(r *server.Hertz, c *config.Config) {
	db, err := mysql.Init(&c.DBConfig)
	if err != nil {
		logs.Errorf("cannot connect to mysql database, err: %v", err)
		//panic(err)
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

	userGroup := r.Group("/user", middleware.Access())
	{
		userHandler := handler.NewUserHandler(supos.NewSuposService(db, &c.SuposConfig))
		userGroup.GET(
			"/list",
			middleware.Response(
				"/user/list",
				userHandler.GetAllUser,
				&iiotpb.GetUserRequest{}))
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
			"/query",
			middleware.Response(
				"/tdengine/query",
				tdengineHandler.Query,
				&iiotpb.TDEngineQueryRequest{}))
	}

	{ // frontend, for debug purpose only
		//r.Static("/web", "deployment/web/html/")

		/*r.StaticFile("/public/web", "public/index.html")
		r.StaticFile("/public/openiiot_white.png", "public/openiiot_white.png")
		r.StaticFile("/public/http.js", "public/http.js")
		r.StaticFile("/public/axios.min.js", "public/axios.min.js")
		r.StaticFile("/public/axios.min.map", "public/axios.min.map")
		r.StaticFile("/public/script.js", "public/script.js")*/

		//fs := &app.FS{Root: "deployment/web/html"}
		//r.StaticFS("/html", fs)
	}
}
