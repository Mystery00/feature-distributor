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

var clientMiddleware gin.HandlerFunc = func(c *gin.Context) {
	if !strings.HasPrefix(c.Request.RequestURI, fmt.Sprintf("%s%s", web.ApiPrefix, "/client")) {
		c.Next()
		return
	}
	clientKey := c.Request.Header.Get("client-key")
	if clientKey == "" {
		resp.Fail(c, 401, "Unauthorized")
		return
	}
	client := grpc.GetCoreClient()
	project, err := client.CheckProject(c.Request.Context(), &pb.CheckProjectRequest{
		ClientKey: &clientKey,
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
