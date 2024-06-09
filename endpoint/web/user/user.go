package user

import (
	"github.com/gin-gonic/gin"
)

func Handle(engine *gin.RouterGroup) {
	engine.POST("/login/v1", login)
	engine.DELETE("/logout/v1", logout)
}
