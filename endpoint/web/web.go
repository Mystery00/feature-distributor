package web

import (
	"feature-distributor/endpoint/web/client"
	"feature-distributor/endpoint/web/health"
	"feature-distributor/endpoint/web/project"
	"feature-distributor/endpoint/web/toggle"
	"feature-distributor/endpoint/web/user"
	"fmt"
	"github.com/gin-gonic/gin"
)

const ApiPrefix = "/api/rest"

func Handle(engine *gin.Engine) {
	client.Handle(engine.Group(fmt.Sprintf("%s%s", ApiPrefix, "/client")))
	user.Handle(engine.Group(fmt.Sprintf("%s%s", ApiPrefix, "/user")))
	health.Handle(engine.Group(fmt.Sprintf("%s%s", ApiPrefix, "/health")))
	toggle.Handle(engine.Group(fmt.Sprintf("%s%s", ApiPrefix, "/toggle")))
	project.Handle(engine.Group(fmt.Sprintf("%s%s", ApiPrefix, "/project")))
}
