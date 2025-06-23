package cli

import (
	"fmt"
	"os"

	"github.com/prodXCE/orbital/runtime"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "orbital",
	Short: "Orbital is a simple, portable container engine.",
	Long: `A learning project to build a lightweight, cross-platform
containerization tool from scratch in Go.`,
}

var runCmd = &cobra.Command{
	Use:   "run [command]",
	Short: "Run a command inside a new container",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		manager, err := runtime.GetManager()
		if err != nil {
			fmt.Printf("Fatal error: could not create container manager: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Printf("--> Using backend of type: %T\n", manager)

		// The 'Start' method now returns a container ID string directly.
		containerID, err := manager.Start(args[0], args[1:])
		if err != nil {
			fmt.Printf("Error starting container: %v\n", err)
			os.Exit(1)
		}

		// We print the ID we received.
		fmt.Printf("Container started successfully! ID: %s\n", containerID)
	},
}

var statusCmd = &cobra.Command{
	Use:   "status [containerID]",
	Short: "Show the status of a container",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		manager, err := runtime.GetManager()
		if err != nil {
			fmt.Printf("Fatal error: could not create container manager: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("--> Using backend of type: %T\n", manager)

		status, err := manager.Status(args[0])
		if err != nil {
			fmt.Printf("Error getting container status: %v\n", err)
			os.Exit(1)
		}
		
		fmt.Printf("Status for container %s: %s\n", args[0], status)
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(statusCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}