package middleware

import "github.com/gin-gonic/gin"

var noRouteMiddleware gin.HandlerFunc = func(context *gin.Context) {
	context.JSON(404, gin.H{
		"message": "404 Not Found",
	})
}
