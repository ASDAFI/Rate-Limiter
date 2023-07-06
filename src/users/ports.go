package users

import (
	"context"
	"time"
)

type IRepository interface {
	Save(ctx context.Context, user *User) error
	FindByUsername(ctx context.Context, username string) (*User, error)
	FindById(ctx context.Context, userId uint) (*User, error)
	FindTokenByUsername(ctx context.Context, username string) (*AuthToken, error)
	IncrementKey(ctx context.Context, key string) error
	IsExistKey(ctx context.Context, key string) (bool, error)
	GetKey(ctx context.Context, key string) (uint, error)
	SetKey(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}
