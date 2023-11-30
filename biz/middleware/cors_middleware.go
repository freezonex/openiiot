package middleware

import (
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/cors"
)

// CorsMiddleware returns a configured Hertz Ext CORS middleware function
// MDN Reference: https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS
func CorsMiddleware() app.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "PATCH", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Length", "Content-Type", "Content-Disposition", "Content-Range",
			"Accept", "X-TT-Access", "X-TT-ENV", "X-USE-BOE", "X-USE-PPE", "X-CUSTOM-TOKEN", "user_token"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	})
}
