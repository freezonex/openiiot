package handler

import (
	"context"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/supos"
	"freezonex/openiiot/biz/service/user"
)

type UserHandler struct {
	suposService *supos.SuposService
	userService  *user.UserService
}

func NewSuposHandler(s *supos.SuposService) *UserHandler {
	return &UserHandler{suposService: s}
}

func NewUserHandler(s *user.UserService) *UserHandler {
	return &UserHandler{userService: s}
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

func (a *UserHandler) GetUser(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetUserRequest)
	resp, err := a.userService.GetUser(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetUser error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *UserHandler) UpdateUser(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.UpdateUserRequest)
	resp, err := a.userService.UpdateUser(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=UpdateUser error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *UserHandler) DeleteUser(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.DeleteUserRequest)
	resp, err := a.userService.DeleteUser(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=DeleteUser error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *UserHandler) AddUser(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.AddUserRequest)
	resp, err := a.userService.AddUser(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AddUser error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *UserHandler) GetUserByToken(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetUserByTokenRequest)
	resp, err := a.userService.GetUserByToken(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=GetUserByToken error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
