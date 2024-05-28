package health

import "github.com/gin-gonic/gin"

func Handle(engine *gin.RouterGroup) {
	engine.GET("/v1", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "ok",
		})
	})
}
