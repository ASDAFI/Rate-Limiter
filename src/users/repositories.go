package users

import (
	"context"
	"github.com/redis/go-redis/v9"
	"service/src/infrastructure/cache"
	"service/src/infrastructure/db"
	"strconv"
	"sync"
	"time"
)

var IncrLock sync.Mutex
var SetLock sync.Mutex

type UserRepository struct {
	dBInfrastructure    db.Provider
	cacheInfrastructure cache.Provider
}

func NewUserRepository(dBInfrastructure db.Provider, cacheInfrastructure cache.Provider) *UserRepository {
	return &UserRepository{dBInfrastructure: dBInfrastructure, cacheInfrastructure: cacheInfrastructure}
}

func (r *UserRepository) Save(ctx context.Context, user *User) error {
	return r.dBInfrastructure.DB.WithContext(ctx).Save(user).Error
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (*User, error) {
	user := &User{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("username = ?", username).Find(user).Error
	return user, err
}

func (r *UserRepository) FindById(ctx context.Context, userId uint) (*User, error) {
	user := &User{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("id = ?", userId).Find(user).Error
	return user, err

}

func (r *UserRepository) FindTokenByUsername(ctx context.Context, username string) (*AuthToken, error) {
	token := &AuthToken{}
	err := r.dBInfrastructure.DB.WithContext(ctx).Where("username = ?", username).Find(token).Error
	return token, err
}

func (r *UserRepository) IncrementKey(ctx context.Context, key string) error {
	IncrLock.Lock()
	defer IncrLock.Unlock()
	err := r.cacheInfrastructure.Client.Incr(ctx, key).Err()
	if err != nil && err != redis.Nil {
		return err
	}
	return nil

}

func (r *UserRepository) IsExistKey(ctx context.Context, key string) (bool, error) {

	exists, err := r.cacheInfrastructure.Client.Exists(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return false, err
	}
	if exists == 0 {
		return false, nil
	}
	return true, nil

}

func (r *UserRepository) GetKey(ctx context.Context, key string) (uint, error) {

	result, err := r.cacheInfrastructure.Client.Get(ctx, key).Result()
	if err != nil && err != redis.Nil {
		return 0, err
	}
	value, err := strconv.ParseUint(result, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(value), nil

}

func (r *UserRepository) SetKey(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	SetLock.Lock()
	defer SetLock.Unlock()
	err := r.cacheInfrastructure.Client.Set(ctx, key, value, expiration).Err()
	if err != nil && err != redis.Nil {
		return err
	}
	return nil

}
