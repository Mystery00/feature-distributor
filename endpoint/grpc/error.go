package grpc

import (
	"context"
	"feature-distributor/common/alert"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"strconv"
)

func errorHandle(ctx context.Context, method string, req, reply any, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	err := invoker(ctx, method, req, reply, cc, opts...)
	if err == nil {
		return nil
	}
	s := status.Convert(err)
	i, err := strconv.Atoi(s.Message())
	if err != nil {
		logrus.Errorf("grpc error: %s", s.Message())
		return err
	}
	return alert.Error(alert.Code(i))
}

func HandleGRPCError(ctx *gin.Context, err error) {
	if c := alert.Convert(err); c != nil {
		e := ReturnErrorMessage(*c)
		ctx.JSON(e.Status, gin.H{"error": e.Msg})
	} else {
		ctx.JSON(500, gin.H{"error": err.Error()})
	}
}
