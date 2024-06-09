package project

import "github.com/gin-gonic/gin"

func Handle(engine *gin.RouterGroup) {
	engine.GET("/list/v1", list)
}
