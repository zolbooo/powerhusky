package core

import (
	"context"
	"os"

	compute "google.golang.org/api/compute/v1"
)

func getInstance(client *compute.Service) (*compute.Instance, error) {
	return client.Instances.Get(os.Getenv(GCP_PROJECT), os.Getenv(GCE_INSTANCE_REGION), os.Getenv(GCE_INSTANCE_ID)).Do()
}

func GetInstanceIP(ctx context.Context, instance *compute.Instance) (string, error) {
	// See: https://cloud.google.com/compute/docs/instances/view-ip-address#api
	for _, iface := range instance.NetworkInterfaces {
		for _, accessConfig := range iface.AccessConfigs {
			if accessConfig.NatIP != "" {
				return accessConfig.NatIP, nil
			}
		}
	}
	return "", nil
}
