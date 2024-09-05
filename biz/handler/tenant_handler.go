package handler

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/tenant"
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

func (a *TenantHandler) GetAllTenantName(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetAllTenantNameRequest)
	resp, err := a.tenantService.GetAllTenantName(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetAllTenantName error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *TenantHandler) AddTenantComponent(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddTenantComponentRequest)
	resp, err := a.tenantService.AddTenantComponent(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddTenantComponent error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *TenantHandler) DeleteTenantComponent(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteTenantComponentRequest)
	resp, err := a.tenantService.DeleteTenantComponent(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteTenantComponent error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *TenantHandler) GetTenantComponent(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetTenantComponentRequest)
	resp, err := a.tenantService.GetTenantComponent(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetTenantComponent error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
