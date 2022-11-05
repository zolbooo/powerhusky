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
	if token != rpc.Token {
		return InvalidToken
	}
	return core.ScheduleShutdown()
}

func (rpc *RPCHandler) PushTask(token string) error {
	if token != rpc.Token {
		return InvalidToken
	}

	counterData, err := core.LoadCounterData(rpc.CounterFile)
	if err != nil {
		return err
	}
	counterData.Counter += 1
	if err = counterData.Save(rpc.CounterFile); err != nil {
		return err
	}

	return nil
}

func (rpc *RPCHandler) RequestShutdown(token string) error {
	if token != rpc.Token {
		return InvalidToken
	}

	counterData, err := core.LoadCounterData(rpc.CounterFile)
	if err != nil {
		return err
	}
	counterData.Counter -= 1
	if err = counterData.Save(rpc.CounterFile); err != nil {
		return err
	}

	if counterData.Counter == 0 {
		core.Shutdown()
	}
	return nil
}
