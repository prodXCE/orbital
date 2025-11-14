package cmd

import (
	"fmt"
	"os"

	"github.com/prodXCE/orbital/downloader"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(pullCmd)
}

var pullCmd = &cobra.Command{
	Use:   "pull <image-name>",
	Short: "Download a rootfs image (e.g., 'alpine-arm' or 'alpine-amd')",
	Long:  `Download and extracts a known rootfs image into the ./orbital/images directory.`,
	Args:  cobra.ExactArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		imageName := args[0]
		fmt.Printf("Pulling image '%s'...\n", imageName)

		if err := downloader.Pull(imageName); err != nil {
			fmt.Printf("Error pulling image: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Image '%s' pulled successfully.\n", imageName)
	},
}
