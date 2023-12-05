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
