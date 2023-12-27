package handler

import (
	"context"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/grafana"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

type GrafanaHandler struct {
	service *grafana.GrafanaService
}

func NewGrafanaHandler(s *grafana.GrafanaService) *GrafanaHandler {
	return &GrafanaHandler{service: s}
}

func (a *GrafanaHandler) Authorize(ctx context.Context, c *app.RequestContext) {
	a.service.Authorize(ctx, c)
}

func (a *GrafanaHandler) GetAccessToken(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GrafanaAccessTokenRequest)
	resp, err := a.service.GetAccessToken(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetAccessToken error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *GrafanaHandler) GetUser(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GrafanaUserRequest)
	resp, err := a.service.GetUser(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetUser error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *GrafanaHandler) GetDatasources(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GrafanaDataSourcesRequest)
	resp, err := a.service.GetDatasources(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetDataSources error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *GrafanaHandler) CreateDatasource(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GrafanaCreateDataSourceRequest)
	resp, err := a.service.NewDataSource(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=CreateDataSource error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *GrafanaHandler) DeleteDatasource(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GrafanaDeleteDataSourceRequest)
	resp, err := a.service.DeleteDataSource(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=CreateDataSource error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *GrafanaHandler) CreateDashBoard(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GrafanaCreateDashboardRequest)
	resp, err := a.service.CreateDashBoard(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=CreateDashboard error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *GrafanaHandler) SaveDashboardByUid(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GrafanaSaveDashboardByUidRequest)
	resp, err := a.service.SaveDashboardByUID(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=SaveDashboardByUid error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}