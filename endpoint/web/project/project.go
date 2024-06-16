package project

import "github.com/gin-gonic/gin"

func Handle(engine *gin.RouterGroup) {
	engine.POST("/v1", create)
	engine.GET("/list/v1", list)
}
