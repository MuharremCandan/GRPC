package gapi

import (
	"test-grpc-project/pkg/model"
	"test-grpc-project/pkg/pb"
)

func converter(user model.User) *pb.User {
	return &pb.User{
		Id:        user.ID,
		Name:      user.Username,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt.String(),
	}
}
