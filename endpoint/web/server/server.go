package server

import "github.com/gin-gonic/gin"

func Handle(engine *gin.RouterGroup) {
	engine.GET("/string/value/v1", stringValue)
}
