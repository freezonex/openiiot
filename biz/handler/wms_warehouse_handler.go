package handler

import (
	"context"
	"freezonex/openiiot/biz/service/wms_warehouse"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

type WmsWarehouseHandler struct {
	wmsWarehouseService *wms_warehouse.WmsWarehouseService
}

func NewWmsWarehouseHandler(s *wms_warehouse.WmsWarehouseService) *WmsWarehouseHandler {
	return &WmsWarehouseHandler{wmsWarehouseService: s}
}

func (a *WmsWarehouseHandler) AddWmsWarehouse(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddWarehouseRequest)
	resp, err := a.wmsWarehouseService.AddWmsWarehouse(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddWmsWarehouse error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsWarehouseHandler) GetWmsWarehouse(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetWarehouseRequest)
	resp, err := a.wmsWarehouseService.GetWmsWarehouse(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetWmsWarehouse error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsWarehouseHandler) UpdateWmsWarehouse(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateWarehouseRequest)
	resp, err := a.wmsWarehouseService.UpdateWmsWarehouse(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateWmsWarehouse error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsWarehouseHandler) DeleteWmsWarehouse(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteWarehouseRequest)
	resp, err := a.wmsWarehouseService.DeleteWmsWarehouse(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteWmsWarehouse error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
