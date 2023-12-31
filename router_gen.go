// Code generated by hertztool

package main

import (
	"freezonex/openiiot/biz/config"
	"freezonex/openiiot/biz/handler"

	"github.com/cloudwego/hertz/pkg/app/server"
)

// register register routes based on the IDL 'api.${HTTP Method}' annotation.
func register(r *server.Hertz, c *config.Config) {
	customizeRegister(r, c)
	r.GET("/ping", handler.Ping)
}
