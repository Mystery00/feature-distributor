package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var noRouteMiddleware gin.HandlerFunc = func(context *gin.Context) {
	logrus.Debugf("No route found: %s", context.Request.RequestURI)
	context.JSON(404, gin.H{
		"message": "404 Not Found",
	})
}
