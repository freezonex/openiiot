package handler

import (
	"context"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/wms_inbound"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

type WmsInboundHandler struct {
	wmsInboundService *wms_inbound.WmsInboundService
}

func NewWmsInboundHandler(s *wms_inbound.WmsInboundService) *WmsInboundHandler {
	return &WmsInboundHandler{wmsInboundService: s}
}

func (a *WmsInboundHandler) AddWmsInbound(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddInboundRequest)
	resp, err := a.wmsInboundService.AddWmsInbound(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddWmsInbound error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsInboundHandler) GetWmsInbound(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetInboundRequest)
	resp, err := a.wmsInboundService.GetWmsInbound(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetWmsInbound error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsInboundHandler) UpdateWmsInbound(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateInboundRequest)
	resp, err := a.wmsInboundService.UpdateWmsInbound(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateWmsInbound error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsInboundHandler) DeleteWmsInbound(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteInboundRequest)
	resp, err := a.wmsInboundService.DeleteWmsInbound(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteWmsInbound error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
