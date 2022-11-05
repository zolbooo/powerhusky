package rpc

import (
	"errors"

	"github.com/zolbooo/powerhusky/daemon/core"
)

var (
	InvalidToken = errors.New("invalid token")
)

type RPCHandler struct {
	Token       string
	CounterFile string
}

func (rpc *RPCHandler) ScheduleShutdown(token string) error {
	if !core.VerifyToken(rpc.Token, token) {
		return InvalidToken
	}
	return core.ScheduleShutdown()
}

func (rpc *RPCHandler) PushTask(token string) error {
	if !core.VerifyToken(rpc.Token, token) {
		return InvalidToken
	}
	_, err := core.EditCounterData(rpc.CounterFile, func(counterData *core.CounterData) {
		counterData.Counter += 1
	})
	return err
}

func (rpc *RPCHandler) RequestShutdown(token string) error {
	if !core.VerifyToken(rpc.Token, token) {
		return InvalidToken
	}

	counterData, err := core.EditCounterData(rpc.CounterFile, func(counterData *core.CounterData) {
		counterData.Counter -= 1
	})
	if err != nil {
		return err
	}

	if counterData.Counter == 0 {
		core.Shutdown()
	}
	return nil
}
