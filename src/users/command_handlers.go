package users

import (
	"context"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type CommandHandler struct {
	UserRepository IRepository
}

func NewCommandHandler(userRepository IRepository) *CommandHandler {
	return &CommandHandler{UserRepository: userRepository}
}

func (h *CommandHandler) CreateUser(ctx context.Context, command CreateUserCommand) (*User, error) {
	// todo : use HashService
	pass, err := bcrypt.GenerateFromPassword([]byte(command.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Info("Password Encryption  failed")
		return nil, err
	}

	user, err := NewUser(CreateUserParameters{
		Username: command.Username,
		Password: string(pass),
		Email:    command.Email,
	})
	if err != nil {
		return nil, err
	}

	err = h.UserRepository.Save(ctx, user)
	return user, err

}

// todo create speific command for this handler
func (h *CommandHandler) RequestRateIncrementByUserIdForCurrentWindow(ctx context.Context, command RequestRateIncrementByUserIdCommand) error {
	currentHash := command.currentTimeHash()
	exists, err := h.UserRepository.IsExistKey(ctx, currentHash)
	if err != nil {
		return err
	}
	if exists {
		err = h.UserRepository.IncrementKey(ctx, currentHash)
	} else {
		err = h.UserRepository.SetKey(ctx, currentHash, 1, time.Minute)
	}
	if err != nil {
		return err // todo: create some errors
	}

	return nil

}

// todo create speific command for this handler
func (h *CommandHandler) RequestRateIncrementByUserIdForLastWindow(ctx context.Context, command RequestRateIncrementByUserIdCommand) error {
	currentHashForLastWindow := command.currentTimeHash() + "last"
	exists, err := h.UserRepository.IsExistKey(ctx, currentHashForLastWindow)
	if err != nil {
		return err
	}
	if exists {
		err = h.UserRepository.IncrementKey(ctx, currentHashForLastWindow)
	} else {
		err = h.UserRepository.SetKey(ctx, currentHashForLastWindow, 1, time.Minute*2)
	}
	if err != nil {
		return err // todo: create some errors
	}

	return nil

}
