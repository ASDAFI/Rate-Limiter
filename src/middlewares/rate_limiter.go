package middlewares

import (
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"service/src/infrastructure/cache"
	"service/src/infrastructure/db"
	"service/src/users"
	"time"
)

func RateLimiterHandler(ctx context.Context) (context.Context, error) {
	method, _ := grpc.Method(ctx)
	userId := ctx.Value("user_id").(uint)

	userRepo := users.NewUserRepository(db.PostgresDBProvider, cache.RedisCacheProvider)
	userQHandler := users.NewUserQueryHandler(userRepo)
	userCHandler := users.NewCommandHandler(userRepo)

	q := users.IsRateLimitedQuery{userId, method, time.Now()}
	r, err := userQHandler.IsRateLimited(ctx, q)
	if err != nil {
		log.Info("Error in rate limiter: ", err)
		return nil, grpc.Errorf(codes.Internal, "some problems")
	}
	if r {
		log.Info("RATE LIMIT!!!!!!")
		return nil, grpc.Errorf(codes.Code(429), "You're limited")
	}

	c := users.UpdateRequestCountCommand{userId, method, time.Now()}
	err = userCHandler.UpdateRequestCount(ctx, c)
	if err != nil {
		log.Info("Error in rate limiter: ", err)
		return nil, grpc.Errorf(codes.Internal, "some problems")
	}
	log.Info("Rate Increment")
	return ctx, nil

}
