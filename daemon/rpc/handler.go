package rpc

import "github.com/zolbooo/powerhusky/daemon/core"

type RPCHandler struct{}

func (rpc *RPCHandler) ScheduleShutdown() error {
	return core.ScheduleShutdown()
}
