package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/wms_material"
)

type WmsMaterialHandler struct {
	wmsMaterialService *wms_material.WmsMaterialService
}

func NewWmsMaterialHandler(s *wms_material.WmsMaterialService) *WmsMaterialHandler {
	return &WmsMaterialHandler{wmsMaterialService: s}
}

func (a *WmsMaterialHandler) AddWmsMaterial(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddMaterialRequest)
	resp, err := a.wmsMaterialService.AddWmsMaterial(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddWmsMaterial error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsMaterialHandler) GetWmsMaterial(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetMaterialRequest)
	resp, err := a.wmsMaterialService.GetWmsMaterial(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetWmsMaterial error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsMaterialHandler) UpdateWmsMaterial(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateMaterialRequest)
	resp, err := a.wmsMaterialService.UpdateWmsMaterial(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateWmsMaterial error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *WmsMaterialHandler) DeleteWmsMaterial(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteMaterialRequest)
	resp, err := a.wmsMaterialService.DeleteWmsMaterial(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteWmsMaterial error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
