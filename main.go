package main

//go get google.golang.org/protobuf/proto
// go get cloud.google.com/go/compute/apiv1
//

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
)

func main() {
	projectID := "myiactool"
	zone := "us-central1-a"
	instanceName := "my-iac-tool-instance"
	machineType := "n1-standard-1"
	sourceImage := "projects/debian-cloud/global/images/debian-12-bookworm-v20241112"
	networkName := "global/networks/default"

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer instancesClient.Close()

	req := &computepb.InsertInstanceRequest{
		Project: projectID,
		Zone:    zone,
		InstanceResource: &computepb.Instance{
			Name: proto.String(instanceName),
			Disks: []*computepb.AttachedDisk{
				{
					InitializeParams: &computepb.AttachedDiskInitializeParams{
						DiskSizeGb:  proto.Int64(10),
						SourceImage: proto.String(sourceImage),
					},
					AutoDelete: proto.Bool(true),
					Boot:       proto.Bool(true),
					Type:       proto.String(computepb.AttachedDisk_PERSISTENT.String()),
				},
			},
			MachineType: proto.String(fmt.Sprintf("zones/%s/machineTypes/%s", zone, machineType)),
			NetworkInterfaces: []*computepb.NetworkInterface{
				{
					Name: proto.String(networkName),
				},
			},
		},
	}

	operation, err := instancesClient.Insert(ctx, req)
	if err != nil {
		log.Fatalf("Failed to create instance: %v", err)
	}

	log.Printf("Instance creation started: %v", operation)
}
