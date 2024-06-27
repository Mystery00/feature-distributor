package toggle

import (
	"feature-distributor/endpoint/grpc"
	"feature-distributor/endpoint/pb"
	"feature-distributor/endpoint/web/resp"
	"github.com/gin-gonic/gin"
)

var remove gin.HandlerFunc = func(c *gin.Context) {
	var t toggleId
	err := c.ShouldBindQuery(&t)
	if err != nil {
		resp.Err(c, 400, err)
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
	resp.Empty(c)
}
