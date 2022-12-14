package core

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/zolbooo/powerhusky/webhook/rpc"
	compute "google.golang.org/api/compute/v1"
)

func StartInstance(ctx context.Context) error {
	client, err := compute.NewService(ctx)
	if err != nil {
		return err
	}

	instance, err := getInstance(client)
	if err != nil {
		return err
	}

	switch instance.Status {
	case "STOPPED":
		fallthrough
	case "TERMINATED":
		op, err := client.Instances.Start(appConfig.GCPProject, appConfig.GCEInstanceRegion, appConfig.GCEInstanceID).Do()
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
		addr, err := GetInstanceIP(ctx, instance)
		if err != nil {
			return err
		}

		closer, rpcClient, err := rpc.CreateClient(ctx, fmt.Sprintf("http://%s:2333", addr))
		if err != nil {
			return err
		}
		defer closer()

		nonce := make([]byte, 16)
		if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
			return err
		}
		return rpcClient.PushTask(GenerateToken(appConfig.DaemonSecret, nonce))
	default:
		return fmt.Errorf("unexpected instance state: %s", instance.Status)
	}

	return nil
}
