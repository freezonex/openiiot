package handler

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/wms_rfid_material"
)

type WmsRfidMaterialHandler struct {
	wmsRfidMaterialService *wms_rfid_material.WmsRfidMaterialService
}

func NewWmsRfidMaterialHandler(s *wms_rfid_material.WmsRfidMaterialService) *WmsRfidMaterialHandler {
	return &WmsRfidMaterialHandler{wmsRfidMaterialService: s}
}

func (a *WmsRfidMaterialHandler) AddWmsRfidMaterial(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddRfidMaterialRequest)
	resp, err := a.wmsRfidMaterialService.AddWmsRfidMaterial(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddWmsRfidMaterial error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsRfidMaterialHandler) GetWmsRfidMaterial(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetRfidMaterialRequest)
	resp, err := a.wmsRfidMaterialService.GetWmsRfidMaterial(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetWmsRfidMaterial error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsRfidMaterialHandler) UpdateWmsRfidMaterial(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateRfidMaterialRequest)
	resp, err := a.wmsRfidMaterialService.UpdateWmsRfidMaterial(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateWmsRfidMaterial error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsRfidMaterialHandler) DeleteWmsRfidMaterial(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteRfidMaterialRequest)
	resp, err := a.wmsRfidMaterialService.DeleteWmsRfidMaterial(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteWmsRfidMaterial error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
