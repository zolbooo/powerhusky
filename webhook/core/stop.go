package core

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/zolbooo/powerhusky/webhook/rpc"
	compute "google.golang.org/api/compute/v1"
)

func StopInstance(ctx context.Context) error {
	client, err := compute.NewService(ctx)
	if err != nil {
		return err
	}

	instance, err := getInstance(client)
	if err != nil {
		return err
	}

	if instance.Status == "RUNNING" {
		addr, err := GetInstanceIP(ctx, instance)
		if err != nil {
			return err
		}

		closer, rpcClient, err := rpc.CreateClient(ctx, fmt.Sprintf("http://%s:2333", addr))
		if err != nil {
			return err
		}
		defer closer()
		return rpcClient.RequestShutdown(os.Getenv(DAEMON_TOKEN))
	}

	log.Printf("Warning: unexpected state when stop was requested - %s", instance.Status)
	return nil
}
