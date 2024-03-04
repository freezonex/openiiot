package middleware

import (
	"context"
	"encoding/json"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"google.golang.org/protobuf/proto"
	"net/http"

	"freezonex/openiiot/biz/model/base_resp"
)

// constant
const (
	HeaderDebugID = "X-RXTX-DebugID"

	metricsAccessKey       = "access"
	metricsHandlerKey      = "handler"
	IDS                    = "IDS"
	REQUEST                = "REQUEST"
	CLIENT                 = "CLIENT"
	HertzMetricesMethodKey = "K_METHOD"
	CHANNEL                = "channel"
)

type HandlerResponse interface {
	proto.Message
	GetBaseResp() *base_resp.BaseResponse
}

var (
	SuccessResponseOK = &base_resp.BaseResponse{
		StatusCode:    http.StatusOK,
		StatusMessage: http.StatusText(http.StatusOK),
	}
	SuccessResponseCreated = &base_resp.BaseResponse{
		StatusCode:    http.StatusCreated,
		StatusMessage: http.StatusText(http.StatusCreated),
	}
	SuccessResponseNoContent = &base_resp.BaseResponse{
		StatusCode:    http.StatusNoContent,
		StatusMessage: http.StatusText(http.StatusNoContent),
	}
)

func ErrorResp(errCode int32, err error) *base_resp.GeneralResponse {
	return &base_resp.GeneralResponse{
		BaseResp: &base_resp.BaseResponse{
			StatusCode:    errCode,
			StatusMessage: err.Error(),
		},
	}
}

type MyHandler func(ctx context.Context, c *app.RequestContext) HandlerResponse

// Access emit request metrics
func Access() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		c.Next(ctx)
	}
}

const (
	mimeJSON     = "application/json"
	mimePROTOBUF = "application/x-protobuf"
)

// pass nil if don't need to parse req
func Response(mkey string, f MyHandler, req proto.Message) app.HandlerFunc {
	handlerFunc := func(ctx context.Context, c *app.RequestContext) {
		c.Set(HertzMetricesMethodKey, mkey)

		var m HandlerResponse
		// Parse the request
		if req != nil {
			newReq := proto.Clone(req)
			if len(c.ContentType()) == 0 {
				logs.CtxWarnf(ctx, "ContentType is empty!")
			}
			//err := binding.BindAndValidate(c, newReq)
			err := c.BindAndValidate(newReq)
			if err != nil {
				m = &base_resp.GeneralResponse{
					BaseResp: &base_resp.BaseResponse{
						StatusCode:    http.StatusBadRequest,
						StatusMessage: err.Error(),
					},
				}
				logs.CtxErrorf(ctx, "Bind request failed, %v", err)
			} else {
				//token := c.GetRequest().Header.Get("X-Custom-Token")
				ctx = context.WithValue(ctx, REQUEST, newReq)

				/*user, err := account.DefaultAccountService().AuthUser(ctx, "", token)
				if err != nil {
					m = &base_resp.GeneralResponse{
						BaseResp: &base_resp.BaseResponse{
							StatusCode:    http.StatusUnauthorized,
							StatusMessage: err.Error(),
						},
					}
					logs.CtxError(ctx, "Error authenticating token")
				} else {
					c.Set("username", user)
					logs.CtxInfo(ctx, "user is %v", user)
				}*/
			}
		}

		// Call the real handler function
		if m == nil {
			m = f(ctx, c)
		}

		// disable any cache, https://stackoverflow.com/questions/49547/how-to-control-web-page-caching-across-all-browsers
		c.Header("Cache-Control", "max-age=0, no-cache, no-store, must-revalidate, proxy-revalidate")
		c.Header("Pragma", "no-cache") // http 1.0
		c.Header("Expires", "0")

		// Write data
		if string(c.ContentType()) == mimePROTOBUF {
			c.ProtoBuf(int(m.GetBaseResp().StatusCode), m)
		} else {
			c.JSON(int(m.GetBaseResp().StatusCode), m)
		}

		// Add log
		var r = ctx.Value(REQUEST)
		if r == nil {
			r = c.Request
		}
		if m.GetBaseResp().StatusCode <= 299 { // means ok
			logs.CtxInfof(ctx, "handler: %s, request: %s, response: %+v", mkey, renderString(r), renderString(m))
		} else if m.GetBaseResp().StatusCode <= 499 { // means client error
			logs.CtxWarnf(ctx, "handler: %s, request: %s, response: %+v", mkey, renderString(r), renderString(m))
		} else { // means server error
			logs.CtxErrorf(ctx, "handler: %s, request: %s, response: %+v", mkey, renderString(r), renderString(m))
		}
		// tagkv["status"] = m.GetBaseResp().StatusCode
		// util.EmitThroughput(metricsHandlerKey, tagkv)
	}
	app.SetHandlerName(handlerFunc, utils.NameOfFunction(f))
	return handlerFunc
}

func renderString(data interface{}) string {
	b, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(b)
}
