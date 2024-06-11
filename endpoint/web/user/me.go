package user

import (
	"github.com/gin-gonic/gin"
)

var me gin.HandlerFunc = func(c *gin.Context) {
	username := c.GetString("username")
	c.JSON(200, gin.H{
		"username": username,
	})
}
