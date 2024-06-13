package project

import (
	"github.com/gin-gonic/gin"
)

var list gin.HandlerFunc = func(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "ok",
		"data": []gin.H{
			{
				"id":   1,
				"name": "Project 1",
			},
			{
				"id":   2,
				"name": "Project 2",
			},
		},
	})
}
