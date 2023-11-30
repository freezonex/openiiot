package handler

import (
	"context"
	"net/http"

	"freezonex/openiiot/biz/middleware"
	"freezonex/openiiot/biz/model/freezonex_openiiot_api"
	"freezonex/openiiot/biz/service/tdengine"

	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
)

type TDEngineHandler struct {
	service *tdengine.TDEngineService
}

func NewTDEngineHandler(s *tdengine.TDEngineService) *TDEngineHandler {
	return &TDEngineHandler{service: s}
}

func (a *TDEngineHandler) TestConnection(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.TDEngineTestConnectionRequest)
	resp, err := a.service.TestConnection(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=TDEngine TestConnection error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *TDEngineHandler) Exec(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.TDEngineExecRequest)
	resp, err := a.service.Exec(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=TDEngine Exec error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}

func (a *TDEngineHandler) Query(ctx context.Context, c *app.RequestContext) middleware.HandlerResponse {
	req := ctx.Value(middleware.REQUEST).(*freezonex_openiiot_api.TDEngineQueryRequest)
	resp, err := a.service.Query(ctx, req, c)
	if err != nil {
		logs.CtxErrorf(ctx, "event=TDEngine Query error=%v", err)
		return middleware.ErrorResp(http.StatusInternalServerError, err)
	}
	return resp
}
