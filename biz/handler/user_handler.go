package handler

import (
	"context"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/supos"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

type UserHandler struct {
	suposService *supos.SuposService
}

func NewUserHandler(s *supos.SuposService) *UserHandler {
	return &UserHandler{suposService: s}
}

func (a *UserHandler) GetSuposUser(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetSuposUserRequest)
	resp, err := a.suposService.GetSuposUser(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetAllUser error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *UserHandler) GetCurrentUser(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetCurrentUserRequest)
	resp, err := a.suposService.GetCurrentUser(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetCurrentUser error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
