package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"service/configs"
	"service/src/infrastructure/db"
	"service/src/users"
)

func Authenticate(ctx context.Context) (context.Context, error) {
	method, ok := grpc.Method(ctx)
	if ok && method == "/service.ritalin.RitalinServer/Login" {
		return ctx, nil
	}
	encodedToken, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	type TokenWithPayload struct {
		users.AuthToken
		UserID uint `json:"user_id"`
	}
	tk := &TokenWithPayload{}
	_, err = jwt.ParseWithClaims(encodedToken, tk, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.Config.Credential.TokenSecret), nil
	})
	err2 := db.PostgresDBProvider.DB.Table(users.AuthToken{}.TableName()).Take(tk).Error
	if err != nil || err2 != nil {
		return nil, grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	ctx = context.WithValue(ctx, "user_id", tk.UserID)

	return ctx, nil
}
