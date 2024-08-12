package middleware

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web"
	"feature-distributor/endpoint/web/resp"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

var serverMiddleware gin.HandlerFunc = func(c *gin.Context) {
	if !strings.HasPrefix(c.Request.RequestURI, fmt.Sprintf("%s%s", web.ApiPrefix, "/server")) {
		c.Next()
		return
	}
	serverKey := c.Request.Header.Get("server-key")
	if serverKey == "" {
		resp.Fail(c, 401, "Unauthorized")
		return
	}
	client := grpc.GetCoreClient()
	project, err := client.CheckProject(c, &pb.CheckProjectRequest{
		ServerKey: &serverKey,
	})
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	if project.GetProject() == nil {
		resp.Fail(c, 401, "Unauthorized")
		return
	}
	c.Set("projectKey", project.GetProject().GetKey())
	c.Next()
}
