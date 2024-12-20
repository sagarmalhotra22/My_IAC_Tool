package cmd

import (
	"github.com/spf13/cobra"
)

// It is the base command which called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "my_iac_tool",
	Short: "A simple tool to create and desctroy GCP compute engine",
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() error {
	// Execute the root command
	return rootCmd.Execute()

}

func init() {
	// Subcommands in root command
	// Add command created the infrastructure
	rootCmd.AddCommand(applyCmd)
	// Destroys the infrastructure
	rootCmd.AddCommand(destroyCmd)
}
