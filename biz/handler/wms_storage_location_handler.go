package handler

import (
	"context"
	"freezonex/openiiot/biz/service/wms_storage_location"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

type WmsStorageLocationHandler struct {
	wmsStorageLocationService *wms_storage_location.WmsStorageLocationService
}

func NewWmsStorageLocationHandler(s *wms_storage_location.WmsStorageLocationService) *WmsStorageLocationHandler {
	return &WmsStorageLocationHandler{wmsStorageLocationService: s}
}

func (a *WmsStorageLocationHandler) AddWmsStorageLocationhouse(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddStorageLocationRequest)
	resp, err := a.wmsStorageLocationService.AddStorageLocation(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddWmsStorageLocation error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsStorageLocationHandler) GetWmsStorageLocationhouse(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetStorageLocationRequest)
	resp, err := a.wmsStorageLocationService.GetStorageLocation(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetWmsStorageLocation error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsStorageLocationHandler) UpdateWmsStorageLocation(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateStorageLocationRequest)
	resp, err := a.wmsStorageLocationService.UpdateStorageLocation(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateWmsStorageLocation error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsStorageLocationHandler) DeleteWmsStorageLocation(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteStorageLocationRequest)
	resp, err := a.wmsStorageLocationService.DeleteStorageLocation(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteWmsStorageLocation error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
