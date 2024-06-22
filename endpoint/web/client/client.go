package client

import "github.com/gin-gonic/gin"

func Handle(engine *gin.RouterGroup) {
	engine.POST("/toggle/v1", toggle)
}
