package rpc

import (
	"errors"

	"github.com/zolbooo/powerhusky/daemon/core"
)

var (
	InvalidToken = errors.New("invalid token")
)

type RPCHandler struct {
	Token string
}

func (rpc *RPCHandler) ScheduleShutdown(token string) error {
	if token != rpc.Token {
		return InvalidToken
	}
	return core.ScheduleShutdown()
}
