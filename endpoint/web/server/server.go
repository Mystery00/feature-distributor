package server

import "github.com/gin-gonic/gin"

func Handle(engine *gin.RouterGroup) {
	engine.GET("/ws", wsHandle)
	engine.POST("/toggle/v1", toggle)
}
