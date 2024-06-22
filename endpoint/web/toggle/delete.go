package toggle

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"github.com/gin-gonic/gin"
)

var remove gin.HandlerFunc = func(c *gin.Context) {
	var t toggleId
	err := c.ShouldBindQuery(&t)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	client := grpc.GetToggleClient()
	_, err = client.DeleteToggle(c.Request.Context(), &pb.GetToggleRequest{
		Id: t.Id,
	})
	if err != nil {
		if err != nil {
			grpc.HandleGRPCError(c, err)
			return
		}
	}
	c.Status(204)
}
