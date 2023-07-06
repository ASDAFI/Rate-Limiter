package users

import (
	"context"
)

type UserQueryHandler struct {
	userRepository IRepository
}

func NewUserQueryHandler(userRepository IRepository) *UserQueryHandler {
	return &UserQueryHandler{
		userRepository: userRepository,
	}
}

func (h *UserQueryHandler) GetUserByUsername(ctx context.Context, query GetUserByUsernameQuery) (*User, error) {
	user, err := h.userRepository.FindByUsername(ctx, query.Username)
	return user, err
}

func (h *UserQueryHandler) GetUserById(ctx context.Context, query GetUserByIdQuery) (*User, error) {
	user, err := h.userRepository.FindById(ctx, query.UserId)
	return user, err
}

func (h *UserQueryHandler) GetUserRequestsCountByUserId(ctx context.Context, query GetUserRequestsCountByUserIdQuery) (uint, error) {
	currentHash := query.currentTimeHash()
	currentRequestsCount := uint(0)
	exists, err := h.userRepository.IsExistKey(ctx, currentHash)
	if err != nil {
		return 0, err
	}
	if exists {
		currentRequestsCount, err = h.userRepository.GetKey(ctx, currentHash)
		if err != nil {
			return 0, err
		}
	}

	lastHash := query.lastTimeHash()
	lastRequestsCount := uint(0)
	exists, err = h.userRepository.IsExistKey(ctx, lastHash)
	if err != nil {
		return 0, err
	}
	if exists {
		lastRequestsCount, err = h.userRepository.GetKey(ctx, lastHash)
		if err != nil {
			return 0, err
		}
	}

	requests := currentRequestsCount + uint((1-float32(query.Time.Unix()%60)/60)*float32(lastRequestsCount))

	return requests, nil
}
