package grpc

import (
	"feature-distributor/common/env"
	"feature-distributor/common/logger"
	"feature-distributor/endpoint/pb"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	rpc "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var connection *rpc.ClientConn
var coreClient pb.CoreServiceClient
var toggleClient pb.ToggleServiceClient
var userClient pb.UserServiceClient

func Init() {
	log := logrus.WithField("source", "grpc")
	opts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}
	address := viper.GetString(env.GrpcAddress)
	conn, err := rpc.NewClient(address,
		rpc.WithTransportCredentials(insecure.NewCredentials()),
		rpc.WithChainUnaryInterceptor(
			logging.UnaryClientInterceptor(logger.InterceptorLogger(log), opts...),
			errorHandle,
		),
	)
	if err != nil {
		panic(err)
	}
	connection = conn
	coreClient = pb.NewCoreServiceClient(conn)
	userClient = pb.NewUserServiceClient(conn)
	toggleClient = pb.NewToggleServiceClient(conn)
}

func Close() error {
	return connection.Close()
}

func GetCoreClient() pb.CoreServiceClient {
	return coreClient
}

func GetUserClient() pb.UserServiceClient {
	return userClient
}

func GetToggleClient() pb.ToggleServiceClient {
	return toggleClient
}
