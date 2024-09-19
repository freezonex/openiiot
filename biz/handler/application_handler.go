package handler

import (
	"context"
	"net/http"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	application "freezonex/openiiot/biz/service/application"
)

type ApplicationHandler struct {
	applicationService *application.ApplicationService
}

func NewApplicationHandler(s *application.ApplicationService) *ApplicationHandler {
	return &ApplicationHandler{applicationService: s}
}

func (a *ApplicationHandler) GetApplication(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetApplicationRequest)
	resp, err := a.applicationService.GetApplication(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetApplication error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

/*func (a *ApplicationHandler) UpdateApplication(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateApplicationRequest)
	resp, err := a.applicationService.UpdateApplication(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateApplication error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}*/

func (a *ApplicationHandler) DeleteApplication(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteApplicationRequest)
	resp, err := a.applicationService.DeleteApplication(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteApplication error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *ApplicationHandler) AddApplication(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddApplicationRequest)
	resp, err := a.applicationService.AddApplication(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddApplication error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *ApplicationHandler) StartApplication(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.StartApplicationRequest)
	resp, err := a.applicationService.StartApplication(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=StartApplication error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *ApplicationHandler) StopApplication(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.StopApplicationRequest)
	resp, err := a.applicationService.StopApplication(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=StopApplication error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
