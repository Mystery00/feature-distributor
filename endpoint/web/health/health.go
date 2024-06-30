package health

import (
	"feature-distributor/endpoint/i18n"
	"github.com/gin-gonic/gin"
)

func Handle(engine *gin.RouterGroup) {
	engine.GET("/v1", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
			"locale":  i18n.Translate(c, "common.health"),
		})
	})
}
