package gapi

import (
	"context"
	"test-grpc-project/pkg/model"
	"test-grpc-project/pkg/pb"
	"test-grpc-project/pkg/utils"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	hashedPassword, err := utils.HashPassword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	user := model.User{
		Username:  req.GetName(),
		Password:  hashedPassword,
		Email:     req.GetEmail(),
		CreatedAt: time.Now(),
	}

	if err := server.db.Create(&user).Error; err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user :%s", err)
	}

	rsp := &pb.CreateUserResponse{
		User:    converter(user),
		Message: "created user",
	}

	return rsp, nil
}
