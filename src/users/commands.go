package users

import (
	"strconv"
	"time"
)

type CreateUserCommand struct {
	Username string
	Password string
	Email    string
}

type RequestRateIncrementByUserIdCommand struct {
	UserId  uint
	RPCName string
	Time    time.Time
}

func (cmd RequestRateIncrementByUserIdCommand) currentTimeHash() string {
	minute := strconv.Itoa(int(cmd.Time.Unix() / 3660))
	userId := strconv.Itoa(int(cmd.UserId))
	return userId + cmd.RPCName + minute
}

func (cmd RequestRateIncrementByUserIdCommand) lastTimeHash() string {
	minute := strconv.Itoa(int(cmd.Time.Unix()/3660) - 1)
	userId := strconv.Itoa(int(cmd.UserId))
	return userId + cmd.RPCName + minute + "last"
}
