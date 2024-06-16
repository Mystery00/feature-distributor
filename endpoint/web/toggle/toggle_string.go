package toggle

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/gin-gonic/gin"
)

var stringValue gin.HandlerFunc = func(c *gin.Context) {
	client := grpc.GetToggleClient()
	response, err := client.GetStringToggle(c.Request.Context(), &pb.StringToggleRequest{
		ReqUser: &pb.ReqUser{
			RolloutKey: "1111",
			Attributes: make(map[string]string),
		},
		ProjectKey:   "123",
		ToggleKey:    "123",
		DefaultValue: "123",
	})
	if err != nil {
		grpc.HandleGRPCError(c, err)
		return
	}
	c.JSON(200, gin.H{"value": response.GetResultValue()})
}
