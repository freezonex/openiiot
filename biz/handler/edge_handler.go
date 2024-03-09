package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/edge"
)

type EdgeHandler struct {
	edgeService *edge.EdgeService
}

func NewEdgeHandler(s *edge.EdgeService) *EdgeHandler {
	return &EdgeHandler{edgeService: s}
}

func (a *EdgeHandler) GetEdge(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetEdgeRequest)
	resp, err := a.edgeService.GetEdge(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetEdge error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *EdgeHandler) UpdateEdge(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateEdgeRequest)
	resp, err := a.edgeService.UpdateEdge(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateEdge error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *EdgeHandler) DeleteEdge(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteEdgeRequest)
	resp, err := a.edgeService.DeleteEdge(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteEdge error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *EdgeHandler) AddEdge(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddEdgeRequest)
	resp, err := a.edgeService.AddEdge(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddEdge error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
