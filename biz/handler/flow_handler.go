package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/flow"
)

type FlowHandler struct {
	flowService *flow.FlowService
}

func NewFlowHandler(s *flow.FlowService) *FlowHandler {
	return &FlowHandler{flowService: s}
}

func (a *FlowHandler) GetFlow(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetFlowRequest)
	resp, err := a.flowService.GetFlow(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetFlow error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *FlowHandler) UpdateFlow(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateFlowRequest)
	resp, err := a.flowService.UpdateFlow(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateFlow error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *FlowHandler) DeleteFlow(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteFlowRequest)
	resp, err := a.flowService.DeleteFlow(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteFlow error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *FlowHandler) AddFlow(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddFlowRequest)
	resp, err := a.flowService.AddFlow(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddFlow error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *FlowHandler) LoadDemoFlow(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.LoadDemoFlowRequest)
	resp, err := a.flowService.LoadDemoFlow(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=LoadDemo error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
