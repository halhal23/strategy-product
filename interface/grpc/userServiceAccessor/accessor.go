package userserviceaccessor

import (
	"context"
	"fmt"

	userPb "github.com/halhal23/strategy-user/pkg/pb"
)

type (
	UserServiceAccessor struct {
		userClient userPb.UserServiceClient
	}

	IUserServiceAccessor interface {
		// CreateUser(ctx context.Context, req *userPb.CreateUserRequest) (*userPb.CreateUserResponse, error)
		ReadUser(ctx context.Context, req *userPb.ReadUserRequest) (*userPb.ReadUserResponse, error)
		// ListUser(ctx context.Context, req *userPb.ListUserRequest) (*userPb.ListUserResponse, error)
	}
)

func (u *UserServiceAccessor) ReadUser(ctx context.Context, req *userPb.ReadUserRequest) (*userPb.ReadUserResponse, error) {
	res, err := u.userClient.ReadUser(ctx, &userPb.ReadUserRequest{Id: req.Id})
	if err != nil {
		fmt.Println(err)
	}
	return &userPb.ReadUserResponse{
		User: &userPb.User{
			Id: res.User.Id,
			Name: res.User.Name,
			Email: res.User.Email,
		}, 
	}, nil
}