package handler

import (
	"context"
	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/emqx"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

type EmqxHandler struct {
	service *emqx.EmqxService
}

func NewEmqxHandler(s *emqx.EmqxService) *EmqxHandler {
	return &EmqxHandler{service: s}
}

func (a *EmqxHandler) GetStatus(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.EmqxGetStatusRequest)
	resp, err := a.service.GetStatus(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetEmqxStatus error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *EmqxHandler) CreateBridge(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.EmqxCreateBridgeRequest)
	resp, err := a.service.CreateBridge(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=CreateEmqxBridge error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp

}

func (a *EmqxHandler) CreateRule(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.EmqxCreateRuleRequest)
	resp, err := a.service.CreateRule(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=CreateEmqxRule error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp

}
