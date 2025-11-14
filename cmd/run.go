package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/prodXCE/orbital/runner"
)

var hostname string

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(
		&hostname,
		"hostname",
		"H",
		"gobox",
		"Hostname for the container",
	)
}

var runCmd = &cobra.Command{
	Use:   "run [rootfs path] [command]",
	Short: "Run a command inside a new container",
	Long:  `Runs a command in a new, isolated container.`,

	Run: func(cmd *cobra.Command, args []string) {

		if len(args) < 2 {
			fmt.Println("Usage: gobox run <rootfs-path> <command> [args...]")
			os.Exit(1)
		}

		rootfsPath := args[0]
		command := args[1:]

		runner.Run(rootfsPath, hostname, command)
	},

	/*
		THE FIX: We must REMOVE (or comment out) this line.
		This will re-enable flag parsing for 'runCmd'.
	*/
	// DisableFlagParsing: true,
}
