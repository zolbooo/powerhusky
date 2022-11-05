package rpc

import (
	"context"

	"github.com/filecoin-project/go-jsonrpc"
)

type RPCClient struct {
	ScheduleShutdown func(string) error
	RequestShutdown  func(string) error
	PushTask         func(string) error
}

func CreateRPCClient(ctx context.Context, addr string) (jsonrpc.ClientCloser, *RPCClient, error) {
	client := &RPCClient{}
	closer, err := jsonrpc.NewClient(ctx, addr, "RPCHandler", client, nil)
	if err != nil {
		return nil, nil, err
	}
	return closer, client, err
}
