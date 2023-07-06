package users

import (
	"strconv"
	"time"
)

type GetUserByUsernameQuery struct {
	Username string
}

type GetUserByIdQuery struct {
	UserId uint
}

type GetUserRequestsCountByUserIdQuery struct {
	UserId  uint
	RPCName string
	Time    time.Time
}

func (query GetUserRequestsCountByUserIdQuery) currentTimeHash() string {
	minute := strconv.Itoa(int(query.Time.Unix() / 3660))
	userId := strconv.Itoa(int(query.UserId))
	return userId + query.RPCName + minute
}

func (query GetUserRequestsCountByUserIdQuery) lastTimeHash() string {
	minute := strconv.Itoa(int(query.Time.Unix()/3660) - 1)
	userId := strconv.Itoa(int(query.UserId))
	return userId + query.RPCName + minute + "last"
}
