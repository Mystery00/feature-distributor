package web

import (
	"feature-distributor/endpoint/web/health"
	"feature-distributor/endpoint/web/server"
	"fmt"
	"github.com/gin-gonic/gin"
)

const apiPrefix = "/api/rest"

func Handle(engine *gin.Engine) {
	health.Handle(engine.Group(fmt.Sprintf("%s%s", apiPrefix, "/health")))
	server.Handle(engine.Group(fmt.Sprintf("%s%s", apiPrefix, "/server")))
}
