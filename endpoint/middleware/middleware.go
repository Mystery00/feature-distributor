package middleware

import "github.com/gin-gonic/gin"

func SetMiddleware(router *gin.Engine) {
	router.Use(corsMiddleware)
	router.Use(clientMiddleware)
	router.Use(serverMiddleware)
	router.Use(adminSessionMiddleware)
	router.NoRoute(noRouteMiddleware)
}
