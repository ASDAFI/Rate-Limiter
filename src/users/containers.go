package users

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username  string `gorm:"column:username;unique_index;not null;"`
	Password  string `gorm:"column:password;not null;"`
	FirstName string `gorm:"column:first_name;"`
	Email     string `gorm:"column:email"`
}

type AuthToken struct {
	Username string `gorm:"column:username"`
	*jwt.StandardClaims
}

type CreateUserParameters struct {
	Username  string
	Password  string
	Email     string
	FirstName string
}

func NewUser(params CreateUserParameters) (*User, error) {
	if err := validatePassword(params.Password); err != nil {
		return nil, err
	}
	user := &User{
		Username:  params.Username,
		Password:  params.Password,
		Email:     params.Email,
		FirstName: params.FirstName,
	}
	return user, user.validateForCreateNewInstance()
}

func (User) TableName() string {
	return "auth_user"
}

func (AuthToken) TableName() string {
	return "auth_token"
}

type TokenWithPayload struct {
	*AuthToken
	UserId uint `json:"user_id"`
}
