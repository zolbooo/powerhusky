package rpc

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/filecoin-project/go-jsonrpc"
)

type TLSOptions struct {
	CertFile string
	KeyFile  string
}

func InitServer(ctx context.Context, rpcHandler *RPCHandler, port int, tlsOptions *TLSOptions) {
	rpcServer := jsonrpc.NewServer()
	rpcServer.Register("RPCHandler", rpcHandler)

	server := &http.Server{Addr: fmt.Sprintf(":%d", port), Handler: rpcServer}
	if tlsOptions != nil {
		go server.ListenAndServeTLS(tlsOptions.CertFile, tlsOptions.KeyFile)
	} else {
		go server.ListenAndServe()
	}

	go func() {
		<-ctx.Done()

		cancelCtx, cancel := context.WithTimeout(context.Background(), time.Second*30)
		server.Shutdown(cancelCtx)
		cancel()
	}()
}
