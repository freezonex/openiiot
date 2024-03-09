package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/wms_outbound"
)

type WmsOutboundHandler struct {
	wmsOutboundService *wms_outbound.WmsOutboundService
}

func NewWmsOutboundHandler(s *wms_outbound.WmsOutboundService) *WmsOutboundHandler {
	return &WmsOutboundHandler{wmsOutboundService: s}
}

func (a *WmsOutboundHandler) AddWmsOutbound(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddOutboundRequest)
	resp, err := a.wmsOutboundService.AddWmsOutbound(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddWmsOutbound error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsOutboundHandler) GetWmsOutbound(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetOutboundRequest)
	resp, err := a.wmsOutboundService.GetWmsOutbound(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetWmsOutbound error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsOutboundHandler) UpdateWmsOutbound(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateOutboundRequest)
	resp, err := a.wmsOutboundService.UpdateWmsOutbound(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateWmsOutbound error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsOutboundHandler) DeleteWmsOutbound(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteOutboundRequest)
	resp, err := a.wmsOutboundService.DeleteWmsOutbound(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteWmsOutbound error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
