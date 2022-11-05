package core

import (
	"context"
	"fmt"
	"os"

	"github.com/zolbooo/powerhusky/webhook/rpc"
	compute "google.golang.org/api/compute/v1"
)

func StartInstance(ctx context.Context) error {
	client, err := compute.NewService(ctx)
	if err != nil {
		return err
	}

	instance, err := getInstance(ctx)
	if err != nil {
		return err
	}

	switch instance.Status {
	case "STOPPED":
		op, err := client.Instances.Start(os.Getenv(GCP_PROJECT), os.Getenv(GCE_INSTANCE_REGION), os.Getenv(GCE_INSTANCE_ID)).Do()
		if err != nil {
			return err
		}
		if op.Error != nil {
			errData, err := op.Error.MarshalJSON()
			if err != nil {
				return fmt.Errorf("failed to encode operation data: %w", err)
			}
			return fmt.Errorf("failed to start instance: %s", string(errData))
		}
	case "RUNNING":
		addr, err := GetInstanceIP(ctx)
		if err != nil {
			return err
		}

		closer, rpcClient, err := rpc.CreateClient(ctx, fmt.Sprintf("http://%s:2333", addr))
		if err != nil {
			return err
		}
		defer closer()
		return rpcClient.PushTask(os.Getenv(DAEMON_TOKEN))
	}

	return nil
}
