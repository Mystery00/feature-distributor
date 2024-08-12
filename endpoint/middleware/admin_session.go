package middleware

import (
	"encoding/json"
	"feature-distributor/endpoint/constants"
	"feature-distributor/endpoint/redis"
	"feature-distributor/endpoint/web"
	"feature-distributor/endpoint/web/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
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
	if strings.HasPrefix(c.Request.RequestURI, fmt.Sprintf("%s%s", web.ApiPrefix, "/client")) {
		c.Next()
		return
	}
	if strings.HasPrefix(c.Request.RequestURI, fmt.Sprintf("%s%s", web.ApiPrefix, "/server")) {
		c.Next()
		return
	}
	authorization := c.Request.Header.Get("Authorization")
	if authorization == "" {
		resp.Fail(c, 401, "Unauthorized")
		return
	}
	key := fmt.Sprintf("session:%s", authorization)
	value, err := redis.Get(c, key)
	if err != nil {
		resp.Err(c, 500, err)
		return
	}
	if value == nil {
		resp.Fail(c, 401, "Unauthorized")
		return
	}
	err = redis.Expire(c, key, constants.UserSessionExpire)
	if err != nil {
		resp.Err(c, 500, err)
		return
	}
	session := make(map[string]any)
	err = json.Unmarshal([]byte(*value), &session)
	if err != nil {
		resp.Err(c, 500, err)
		return
	}
	c.Set("userId", session["userId"])
	c.Set("username", session["username"])
	c.Next()
}
