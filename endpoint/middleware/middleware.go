package middleware

import "github.com/gin-gonic/gin"

func SetMiddleware(router *gin.Engine) {
	router.Use(adminSessionMiddleware)
	router.NoRoute(noRouteMiddleware)
}
