package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/sagarmalhotra22/My_IAC_Tool/state"

	compute "cloud.google.com/go/compute/apiv1"
	"github.com/spf13/cobra"
	computepb "google.golang.org/genproto/googleapis/cloud/compute/v1"
)

// DestroyCmd represents the destroy command to delete a VM instance.
var DestroyCmd = &cobra.Command{
	Use:   "destroy",
	Short: "Delete the GCP VM instance",
	Run: func(cmd *cobra.Command, args []string) {
		// Loading state file
		stateData, err := state.LoadState()
		if err != nil {
			log.Fatalf("Failed to load state: %v", err)
		}

		if stateData == nil {
			fmt.Println("No instance found in state.json")
			return
		}
		// seting context for the request
		ctx := context.Background()
		instancesClient, err := compute.NewInstancesRESTClient(ctx)
		if err != nil {
			log.Fatalf("Failed to create client: %v", err)
		}
		defer instancesClient.Close()
		// createing delete request
		req := &computepb.DeleteInstanceRequest{
			Project:  stateData.ProjectID,
			Zone:     stateData.Zone,
			Instance: stateData.InstanceName,
		}
		// deleting the instance
		op, err := instancesClient.Delete(ctx, req)
		if err != nil {
			log.Fatalf("Failed to delete instance: %v", err)
		}

		fmt.Printf("VM instance %s deleted. Operation status: %s\n", stateData.InstanceName, op.Status)

		// Removing the state file
		if err := state.DeleteState(); err != nil {
			log.Fatalf("Failed to delete state file: %v", err)
		}
		fmt.Println("State file deleted")
	},
}
