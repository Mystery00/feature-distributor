package user

import "github.com/gin-gonic/gin"

var logout gin.HandlerFunc = func(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "ok",
	})
}
