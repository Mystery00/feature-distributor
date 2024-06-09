package middleware

import (
	"feature-distributor/endpoint/constants"
	"feature-distributor/endpoint/redis"
	"feature-distributor/endpoint/web"
	"fmt"
	"github.com/gin-gonic/gin"
)

var whiteUri = []string{
	web.ApiPrefix + "/health/v1",
	web.ApiPrefix + "/user/login/v1",
}

func isWhiteUri(uri string) bool {
	for _, v := range whiteUri {
		if uri == v {
			return true
		}
	}
	return false
}

var adminSessionMiddleware gin.HandlerFunc = func(context *gin.Context) {
	if isWhiteUri(context.Request.RequestURI) {
		context.Next()
		return
	}
	authorization := context.Request.Header.Get("Authorization")
	if authorization == "" {
		context.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		context.Abort()
		return
	}
	key := fmt.Sprintf("session:%s", authorization)
	value, err := redis.Get(context.Request.Context(), key)
	if err != nil {
		context.JSON(500, gin.H{
			"message": err.Error(),
		})
		context.Abort()
		return
	}
	if value == nil {
		context.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		context.Abort()
		return
	}
	err = redis.Expire(context.Request.Context(), key, constants.UserSessionExpire)
	if err != nil {
		context.JSON(500, gin.H{
			"message": err.Error(),
		})
		context.Abort()
		return
	}
	context.Next()
}
