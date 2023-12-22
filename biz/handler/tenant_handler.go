package handler

import (
	"context"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/tenant"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

type TenantHandler struct {
	tenantService *tenant.TenantService
}

func NewTenantHandler(s *tenant.TenantService) *TenantHandler {
	return &TenantHandler{tenantService: s}
}

func (a *TenantHandler) AddTenant(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddTenantRequest)
	resp, err := a.tenantService.AddTenant(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddTenant error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *TenantHandler) GetTenant(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetTenantRequest)
	resp, err := a.tenantService.GetTenant(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetTenant error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *TenantHandler) UpdateTenant(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateTenantRequest)
	resp, err := a.tenantService.UpdateTenant(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateTenant error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *TenantHandler) DeleteTenant(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteTenantRequest)
	resp, err := a.tenantService.DeleteTenant(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteTenant error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
