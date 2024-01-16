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

type AuthHandler struct {
	service *supos.SuposService
}

func NewAuthHandler(s *supos.SuposService) *AuthHandler {
	return &AuthHandler{service: s}
}

type RestResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func NewRestResult(code int, msg string, data interface{}) *RestResult {
	return &RestResult{
		Code: code,
		Msg:  msg,
		Data: data,
	}
}

func (a *AuthHandler) Authorize(ctx context.Context, c *app.RequestContext) {
	redirectURL := c.GetString("redirect_url")
	location := a.service.Authorize(redirectURL, "1")

	logs.CtxInfof(ctx, "Redirecting to: %s", location)
	c.Response.Header.Set("Location", location)
	c.Response.SetStatusCode(http.StatusFound)
}

func (a *AuthHandler) AccessToken(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.GetAccessTokenRequest)
	resp, err := a.service.GetAccessToken(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=AccessToken error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *AuthHandler) Logout(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.LogoutRequest)
	resp, err := a.service.Logout(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=Logout error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *AuthHandler) Login(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.LoginRequest)
	resp, err := a.service.Login(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=Login error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
