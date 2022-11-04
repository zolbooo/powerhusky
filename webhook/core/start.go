package core

import (
	"context"
	"fmt"
	"os"

	compute "google.golang.org/api/compute/v1"
)

func StartInstance(ctx context.Context) error {
	client, err := compute.NewService(ctx)
	if err != nil {
		return err
	}

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

	return nil
}
