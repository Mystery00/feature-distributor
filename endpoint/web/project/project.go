package project

import "github.com/gin-gonic/gin"

func Handle(engine *gin.RouterGroup) {
	engine.POST("/v1", save)
	engine.GET("/list/v1", list)
	engine.GET("/v1", get)
	engine.DELETE("/v1", remove)
}

type projectId struct {
	Id int64 `form:"id" required:"true" binding:"required"`
}
