package cmd

import (
	"context"
	"fmt"
	"log"

	compute "cloud.google.com/go/compute/apiv1"
	computepb "cloud.google.com/go/compute/apiv1/computepb"

	"github.com/sagarmalhotra22/My_IAC_Tool/config"
	"github.com/sagarmalhotra22/My_IAC_Tool/state"
	"github.com/spf13/cobra"
	"google.golang.org/protobuf/proto"
)

// applyCmd represents the apply command
var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Create a new VM instance",
	Run: func(cmd *cobra.Command, args []string) {
		// Load the current state

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
		// Saving state to state file
		state_obj := &state.State{
			Instances: cfg.InstanceName,
			ProjectId: cfg.ProjectID,
			Zone:      cfg.Zone,
		}

		err = state.SaveState(state_obj)
		if err != nil {
			log.Fatalf("Failed to add state configuration in the state file: %v", err)

		}

		fmt.Printf("VM '%s' created and state saved.\n", cfg.InstanceName)
	},
}
