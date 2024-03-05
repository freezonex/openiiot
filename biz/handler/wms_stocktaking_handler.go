package handler

import (
	"context"
	"freezonex/openiiot/biz/service/wms_stocktaking"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

type WmsStocktakingHandler struct {
	wmsStocktakingService *wms_stocktaking.WmsStocktakingService
}

func NewWmsStocktakingHandler(s *wms_stocktaking.WmsStocktakingService) *WmsStocktakingHandler {
	return &WmsStocktakingHandler{wmsStocktakingService: s}
}

func (a *WmsStocktakingHandler) AddWmsStocktaking(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddStocktakingRequest)
	resp, err := a.wmsStocktakingService.AddStocktaking(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddWmsStocktaking error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsStocktakingHandler) GetWmsStocktaking(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetStocktakingRequest)
	resp, err := a.wmsStocktakingService.GetStocktaking(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetWmsStocktaking error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsStocktakingHandler) UpdateWmsStocktaking(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateStocktakingRequest)
	resp, err := a.wmsStocktakingService.UpdateStocktaking(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateWmsStocktaking error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsStocktakingHandler) DeleteWmsStocktaking(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteStocktakingRequest)
	resp, err := a.wmsStocktakingService.DeleteStocktaking(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteWmsStocktaking error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
