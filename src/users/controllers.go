package users

import (
	"context"
	"github.com/golang/protobuf/ptypes/empty"
	log "github.com/sirupsen/logrus"
	"google.golang.org/protobuf/types/known/emptypb"
	"service/src/infrastructure/cache"
	"service/src/infrastructure/db"
	pb_user "service/src/proto/messages/user"
)

type UserServer struct{}

func (s UserServer) Login(ctx context.Context, request *pb_user.LoginRequest) (*pb_user.LoginResponse, error) {
	log.Info("Login -- username: ", request.Username)
	userRepo := NewUserRepository(db.PostgresDBProvider, cache.RedisCacheProvider)
	authHandler := NewAuthHandler(userRepo)
	loginErr, token := authHandler.Login(ctx, request.GetUsername(), request.GetPassword())
	return &pb_user.LoginResponse{
		Token: token,
	}, loginErr
}

func (s UserServer) GetUser(ctx context.Context, empty *emptypb.Empty) (*pb_user.User, error) {

	userId := ctx.Value("user_id").(uint)

	log.Info("GetUser -- userId: ", userId)

	userRepo := NewUserRepository(db.PostgresDBProvider, cache.RedisCacheProvider)
	userQHandler := NewUserQueryHandler(userRepo)

	log.Info("Rate Increment")

	query := GetUserByIdQuery{UserId: userId}
	user, err := userQHandler.GetUserById(ctx, query)
	if err != nil {
		return nil, err
	}
	return &pb_user.User{
		Username: user.Username,
		Email:    user.Email,
	}, nil

}

func (s UserServer) Logout(ctx context.Context, nothing *emptypb.Empty) (*empty.Empty, error) {
	userId := ctx.Value("user_id").(uint)
	log.Info("Logout -- userId: ", userId)

	userRepo := NewUserRepository(db.PostgresDBProvider, cache.RedisCacheProvider)
	authHandler := NewAuthHandler(userRepo)
	err := authHandler.Logout(ctx, userId)

	return &empty.Empty{}, err
}
