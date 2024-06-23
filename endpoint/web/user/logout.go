package user

import (
	"feature-distributor/endpoint/redis"
	"feature-distributor/endpoint/web/resp"
	"fmt"
	"github.com/gin-gonic/gin"
)

var logout gin.HandlerFunc = func(c *gin.Context) {
	token := c.GetHeader("Authorization")
	key := fmt.Sprintf("session:%s", token)
	err := redis.Del(c.Request.Context(), key)
	if err != nil {
		resp.Err(c, 500, err)
		return
	}
	resp.Empty(c)
}
