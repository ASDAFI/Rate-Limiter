package users

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"service/configs"
	"service/src/infrastructure/db"
	"time"
)

type AuthHandler struct {
	userRepo IRepository
}

func NewAuthHandler(userRepo IRepository) *AuthHandler {
	return &AuthHandler{userRepo: userRepo}
}

// todo : separate functions
func (h *AuthHandler) Login(ctx context.Context, username string, password string) (error, string) {
	user, err := h.userRepo.FindByUsername(ctx, username)
	if err != nil {
		return err, ""
	}

	expiresAt := time.Now().Add(time.Minute * 100000).Unix()
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return grpc.Errorf(codes.Unauthenticated, "invalid auth token: %v", err), ""
	}

	tk := &AuthToken{
		Username: user.Username,
		StandardClaims: &jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Id:        fmt.Sprintf("%s", uuid.NewV4()),
		},
	}

	tokenWithPayload := TokenWithPayload{
		AuthToken: tk,
		UserId:    user.ID,
	}
	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), tokenWithPayload)
	var tokenString string
	tokenString, err = token.SignedString([]byte(configs.Config.Credential.TokenSecret))
	if err != nil {
		log.Info(err)
	}
	createToken := db.PostgresDBProvider.DB.Create(tk)

	if createToken.Error != nil {
		log.Info(createToken.Error)
	}
	return nil, tokenString
}

func (h *AuthHandler) Logout(ctx context.Context, userId uint) error {
	user, err := h.userRepo.FindById(ctx, userId)
	if err != nil {
		return err
	}
	token, err := h.userRepo.FindTokenByUsername(ctx, user.Username)
	if err != nil {
		return err
	}
	err = db.PostgresDBProvider.DB.WithContext(ctx).Delete(token).Error
	if err != nil {
		return err
	}

	return nil
}
