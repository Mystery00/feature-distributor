package grpc

import (
	"feature-distributor/core/pb"
	"google.golang.org/grpc"
	"net"
)

func Run(addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterCoreServiceServer(server, &CoreServer{})
	pb.RegisterToggleServiceServer(server, &ToggleServer{})
	return server.Serve(lis)
}
