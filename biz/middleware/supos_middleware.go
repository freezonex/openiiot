package middleware

import (
	"context"
	"freezonex/openiiot/biz/service/callback_mgr"
	"net/http"
	"regexp"
	"strings"
	"time"

	"freezonex/openiiot/biz/model/base_resp"
	"github.com/cloudwego/hertz/pkg/app"
	logs "github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"google.golang.org/protobuf/proto"
)

const (
	userTokenHeader = "userToken" // This is the name of the header where the token is expected to be
	whiteList       = "/auth/*;/ping;/grafana/*;/emqx/*;/tdengine;/public;/tenant/allname"
)

func AuthMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// Check if the request path is in the white list
		// Assume `isWhiteListed` is a function you should implement that checks the white list
		logs.CtxDebugf(ctx, "request URI: %v", string(c.Request.URI().Path()))
		if isWhiteListed(string(c.Request.URI().Path())) {
			c.Next(ctx)
			return
		}

		// Get the userToken from the header
		userToken := string(c.Request.Header.Peek(userTokenHeader))

		// Check if the userToken is not blank and is in the database
		if userToken != "" {
			// Check the token
			foundUpdatetime, foundUser, err := CallGetUserByTokenDB(ctx, userToken)

			if err != nil {
				c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
					"code": http.StatusInternalServerError,
					"msg":  "Internal server error",
				})
				return
			}

			//current time - token update time <= 1200
			if foundUser != "" && foundUpdatetime != nil && time.Since(*foundUpdatetime) <= 1200*time.Second {
				// If the token exists and is valid, continue with the next middleware or handler
				// Add username to context and continue
				ctxWithUsername := context.WithValue(ctx, "currentusername", foundUser)
				c.Next(ctxWithUsername)
				return
			}
		}

		// If we got here, the user is not authorized, so we send a 401 response
		c.AbortWithStatusJSON(http.StatusUnauthorized, map[string]interface{}{
			"code": http.StatusUnauthorized,
			"msg":  "Unauthorized access",
		})

		// If the token is not present or not valid, stop the request and respond with 401
		/*c.Response.WriteHeader(http.StatusUnauthorized)
		c.JSON(http.StatusUnauthorized, ErrorResp(http.StatusUnauthorized, fmt.Errorf("unauthorized: no permissions to access this resource")))
		// Stop the chain of handlers/middlewares here, because the user is not authorized
		c.Abort()*/
	}
}

func isWhiteListed(uri string) bool {
	excludeUris := strings.Split(whiteList, ";")
	for _, pattern := range excludeUris {
		if pattern == "" {
			continue // Skip empty patterns
		}

		matched, err := regexp.MatchString(pattern, uri)
		if err != nil {
			// Handle the error according to your application's needs.
			// For example, log it and consider the URI as not whitelisted:
			// log.Printf("Error compiling pattern '%s': %v", pattern, err)
			continue // For now, let's simply continue checking the next pattern.
		}

		if matched {
			return true // The URI matches one of the patterns in the whitelist.
		}
	}
	return false // No match found, the URI is not in the whitelist.
}

// pass nil if don't need to parse req
func AuthResponse(mkey string, f MyHandler, req proto.Message) app.HandlerFunc {
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
				logs.CtxErrorf(ctx, "Bind request failed: %v", err)
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

// use callback to call user's function
func CallGetUserByTokenDB(ctx context.Context, usertoken string) (*time.Time, string, error) {
	updatetime, username, err := callback_mgr.CallUserByTokenDBFunc("GetUserByTokenDB", ctx, usertoken)
	if err != nil {
		return nil, "", err
	}
	// Use reflection to call the function dynamically

	return updatetime, username, err
}
