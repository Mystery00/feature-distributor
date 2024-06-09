package project

import "github.com/gin-gonic/gin"

var list gin.HandlerFunc = func(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "ok",
		"data":    []string{"project1", "project2"},
	})
}
