package cli

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/prodXCE/orbital/runtime" // custom runtime package
)

var rootCmd = &cobra.Command {
	Use: "orbital",
	Short: "Orbital is a simple, portable container engine",
	Long: `A learning project to build a lightweight, cross-platform containerization tool from sratch in Go.`,

}

// runCmd represents the "run" command
var runCmd = &cobra.Command {
	Use: "run [command]",
	Short: "Run a command inside a new container",
	Long: `Run a command inside a new container.`,

	// This makes sure that the user provides at least one argument (the command to run)
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		command := args[0]
		commandArgs := args[1:]
		runtime.RunContainer(command, commandArgs)
	},
}

// statusCmd represents the "status" command
var statusCmd = &cobra.Command {
	Use: "status [containerID]",
	Short: "Show the status of a container",
	Long: `Show the status of a container.`,

	// This command requires exactly one argument: the container ID
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		containerID := args[0]
		fmt.Printf("Checking status for container: %s...\n", containerID)
		fmt.Printf("[INFO] Status command is not implemented yet.")
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
