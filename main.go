package main

//go get google.golang.org/protobuf/proto
// go get cloud.google.com/go/compute/apiv1
//go get gopkg.in/yaml.v3

import (
	"context"
	"fmt"
	"gcpprovision/config"
	"log"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"
	"google.golang.org/protobuf/proto"
)

func main() {
	// Passing config file to the parse config function
	cfg, err := config.ParseConfig("config.yaml")
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	ctx := context.Background()
	instancesClient, err := compute.NewInstancesRESTClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer instancesClient.Close()

	req := &computepb.InsertInstanceRequest{
		Project: cfg.ProjectID,
		Zone:    cfg.Zone,
		InstanceResource: &computepb.Instance{
			Name: proto.String(cfg.InstanceName),
			Disks: []*computepb.AttachedDisk{
				{
					InitializeParams: &computepb.AttachedDiskInitializeParams{
						DiskSizeGb:  proto.Int64(cfg.DiskSizeGb),
						SourceImage: proto.String(cfg.Image),
					},
					AutoDelete: proto.Bool(true),
					Boot:       proto.Bool(true),
					Type:       proto.String(computepb.AttachedDisk_PERSISTENT.String()),
				},
			},
			MachineType: proto.String(fmt.Sprintf("zones/%s/machineTypes/%s", cfg.Zone, cfg.MachineType)),
			NetworkInterfaces: []*computepb.NetworkInterface{
				{
					Name: proto.String(cfg.NetworkName),
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
