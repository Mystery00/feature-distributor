package grpc

import (
	"feature-distributor/common/env"
	"feature-distributor/endpoint/pb"
	"github.com/spf13/viper"
	rpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var connection *rpc.ClientConn
var toggleClient pb.ToggleServiceClient
var userClient pb.UserServiceClient

func Init() {
	address := viper.GetString(env.GrpcAddress)
	conn, err := rpc.NewClient(address, rpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	connection = conn
	userClient = pb.NewUserServiceClient(conn)
	toggleClient = pb.NewToggleServiceClient(conn)
}

func Close() error {
	return connection.Close()
}

func GetUserClient() pb.UserServiceClient {
	return userClient
}

func GetToggleClient() pb.ToggleServiceClient {
	return toggleClient
}
