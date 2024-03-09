package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/core"
)

type CoreHandler struct {
	coreService *core.CoreService
}

func NewCoreHandler(s *core.CoreService) *CoreHandler {
	return &CoreHandler{coreService: s}
}

func (a *CoreHandler) GetCore(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetCoreRequest)
	resp, err := a.coreService.GetCore(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetCore error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *CoreHandler) UpdateCore(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateCoreRequest)
	resp, err := a.coreService.UpdateCore(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateCore error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *CoreHandler) DeleteCore(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteCoreRequest)
	resp, err := a.coreService.DeleteCore(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteCore error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *CoreHandler) AddCore(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddCoreRequest)
	resp, err := a.coreService.AddCore(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddCore error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
