package grpc

import (
	"context"
	"errors"
	"feature-distributor/core/db/model"
	"feature-distributor/core/db/query"
	"feature-distributor/core/pb"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

type UserServer struct {
	pb.UnimplementedUserServiceServer
}

func (u UserServer) CheckLogin(ctx context.Context, in *pb.CheckLoginRequest) (*pb.CheckLoginResponse, error) {
	q := query.User
	qc := q.WithContext(ctx)
	user, err := qc.Where(q.Username.Eq(in.GetUsername())).First()
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			logrus.Infof("user not found: %v", in.GetUsername())
			return &pb.CheckLoginResponse{
				Code: http.StatusUnauthorized,
			}, nil
		}
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(in.GetPassword()))
	if err != nil {
		return &pb.CheckLoginResponse{
			Code: http.StatusUnauthorized,
		}, nil
	}
	return &pb.CheckLoginResponse{
		Code:   http.StatusOK,
		UserId: &user.ID,
	}, nil
}

func (u UserServer) InitAdmin(ctx context.Context, in *pb.InitAdminRequest) (*pb.InitAdminResponse, error) {
	q := query.User
	qc := q.WithContext(ctx)
	user, err := qc.Where(q.Username.Eq("admin")).First()
	if err == nil && user != nil {
		logrus.Info("admin already exists")
		return &pb.InitAdminResponse{
			Code:    http.StatusNotAcceptable,
			Message: "admin user already exists",
		}, nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(in.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	err = q.Create(&model.User{
		Username: "admin",
		Password: string(hash),
	})
	if err != nil {
		return nil, err
	}
	return &pb.InitAdminResponse{
		Code:    http.StatusOK,
		Message: "",
	}, nil
}
