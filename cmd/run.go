package cmd

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/prodXCE/orbital/downloader"
    "github.com/prodXCE/orbital/runner"
)

var hostname string

func init() {
    rootCmd.AddCommand(runCmd)
    runCmd.Flags().StringVarP(&hostname, "hostname", "H", "orbital", "Hostname for the container")
}

var runCmd = &cobra.Command{
    Use:   "run [image-or-path] [command]",
    Short: "Run a command inside a new container",
    Long:  `Runs a command in a new, isolated container.`,

    Run: func(cmd *cobra.Command, args []string) {
        if len(args) < 2 {
            fmt.Println("Usage: orbital run <image-or-path> <command> [args...]")
            os.Exit(1)
        }

        imageOrPath := args[0]
        command := args[1:]

        var rootfsPath string

        if _, err := os.Stat(imageOrPath); err == nil {
            rootfsPath = imageOrPath
            fmt.Printf("Using local path: %s\n", rootfsPath)

        } else if os.IsNotExist(err) {
            fmt.Printf("Local path not found. Checking for downloaded image: %s\n", imageOrPath)

            path, exists := downloader.GetImagePath(imageOrPath)

            if !exists {
                fmt.Printf("Error: Image '%s' not found locally.\n", imageOrPath)
                fmt.Printf("Please run: ./orbital pull %s\n", imageOrPath)
                os.Exit(1)
            }

            rootfsPath = path
            fmt.Printf("Using downloaded image at: %s\n", rootfsPath)

        } else {
            fmt.Printf("Error checking path %s: %v\n", imageOrPath, err)
            os.Exit(1)
        }

        runner.Run(rootfsPath, hostname, command)
    },
}

