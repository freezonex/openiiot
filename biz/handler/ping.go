// Code generated by hertztool.

package handler

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

// Ping .
func Ping(ctx context.Context, c *app.RequestContext) {
	hlog.CtxInfof(ctx, "ping received")
	c.JSON(200, utils.H{
		"message": "pong",
	})
}
