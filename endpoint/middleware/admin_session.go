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

var adminSessionMiddleware gin.HandlerFunc = func(c *gin.Context) {
	if isWhiteUri(c.Request.RequestURI) {
		c.Next()
		return
	}
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	key := fmt.Sprintf("session:%s", authorization)
	value, err := redis.Get(c.Request.Context(), key)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	if value == nil {
		c.JSON(401, gin.H{
			"message": "Unauthorized",
		})
		c.Abort()
		return
	}
	err = redis.Expire(c.Request.Context(), key, constants.UserSessionExpire)
	if err != nil {
		c.JSON(500, gin.H{
			"message": err.Error(),
		})
		c.Abort()
		return
	}
	c.Set("username", *value)
	c.Next()
}
