package web

import (
	"feature-distributor/endpoint/web/health"
	"feature-distributor/endpoint/web/project"
	"feature-distributor/endpoint/web/server"
	"feature-distributor/endpoint/web/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

const ApiPrefix = "/api/rest"

func Handle(engine *gin.Engine) {
	user.Handle(engine.Group(fmt.Sprintf("%s%s", ApiPrefix, "/user")))
	health.Handle(engine.Group(fmt.Sprintf("%s%s", ApiPrefix, "/health")))
	server.Handle(engine.Group(fmt.Sprintf("%s%s", ApiPrefix, "/server")))
	project.Handle(engine.Group(fmt.Sprintf("%s%s", ApiPrefix, "/project")))
}
