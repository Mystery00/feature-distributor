package middleware

import (
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SetMiddleware(router *gin.Engine) {
	router.Use(corsMiddleware)
	router.Use(clientMiddleware)
	router.Use(serverMiddleware)
	router.Use(adminSessionMiddleware)
	router.Use(func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			resp.Err(c, http.StatusInternalServerError, err)
		}
	})
	router.NoRoute(noRouteMiddleware)
}
