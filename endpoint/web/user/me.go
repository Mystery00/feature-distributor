package user

import (
	"feature-distributor/endpoint/web/resp"
	"fmt"
	"github.com/gin-gonic/gin"
)

var me gin.HandlerFunc = func(c *gin.Context) {
	userId := c.GetInt64("userId")
	username := c.GetString("username")
	resp.Data(c, gin.H{
		"userId":   fmt.Sprintf("%d", userId),
		"username": username,
	})
}
