package users

import (
	"strconv"
	"time"
)

type CreateUserCommand struct {
	Username  string
	Password  string
	Email     string
	FirstName string
}

type RequestRateIncrementByUserIdCommand struct {
	UserId  uint
	RPCName string
	Time    time.Time
}

type UpdateRequestCountCommand struct {
	UserId  uint
	RPCName string
	Time    time.Time
}

type SetRateLimitCommand struct {
	RequestsPerMinute uint
	RPCName           string
}

func (cmd RequestRateIncrementByUserIdCommand) currentTimeHash() string {
	minute := strconv.Itoa(int(cmd.Time.Unix() / 60))
	userId := strconv.Itoa(int(cmd.UserId))
	return userId + cmd.RPCName + minute
}

func (cmd RequestRateIncrementByUserIdCommand) lastTimeHash() string {
	minute := strconv.Itoa(int(cmd.Time.Unix()/60) - 1)
	userId := strconv.Itoa(int(cmd.UserId))
	return userId + cmd.RPCName + minute + "last"
}
