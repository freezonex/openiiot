package handler

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/application"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"
)

type AppHandler struct {
	appService *application.AppService
}

func NewAppHandler(s *application.AppService) *AppHandler {
	return &AppHandler{appService: s}
}

func (a *AppHandler) GetApp(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetAppRequest)
	resp, err := a.appService.GetApp(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetApp error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *AppHandler) UpdateApp(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateAppRequest)
	resp, err := a.appService.UpdateApp(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateApp error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *AppHandler) DeleteApp(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteAppRequest)
	resp, err := a.appService.DeleteApp(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteApp error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *AppHandler) AddApp(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddAppRequest)
	resp, err := a.appService.AddApp(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddApp error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
