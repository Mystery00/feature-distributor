package grpc

import (
	"feature-distributor/common/logger"
	"feature-distributor/core/pb"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func Run(addr string) error {
	log := logrus.WithField("source", "grpc")
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	opts := []logging.Option{
		logging.WithLogOnEvents(),
	}
	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(logger.InterceptorLogger(log), opts...),
		),
	)
	pb.RegisterEventServiceServer(server, &EventServer{})
	pb.RegisterHealthServiceServer(server, &HealthServer{})
	pb.RegisterUserServiceServer(server, &UserServer{})
	pb.RegisterCoreServiceServer(server, &CoreServer{})
	pb.RegisterToggleServiceServer(server, &ToggleServer{})
	return server.Serve(lis)
}
