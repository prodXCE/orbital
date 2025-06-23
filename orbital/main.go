package main

import (
	"fmt"
	"os"

	"github.com/prodXCE/orbital/backends"
	"github.com/prodXCE/orbital/cli"
)

func main() {
	// This is the check that separates the parent from the child process.
	// If the program is run with "child" as the first argument, it will
	// execute the container's internal setup instead of the CLI.
	if len(os.Args) > 1 && os.Args[1] == "child" {
		fmt.Println("--> Entering child process execution mode...")
		// We are inside the container, execute the child process logic.
		// This must be gated by the OS, as ChildProcess is in a linux-only file.
		if backends.IsLinux() { // We will add this helper function.
			backends.ChildProcess(os.Args[2], os.Args[3:])
		} else {
			fmt.Println("Error: Child process mode is only supported on Linux.")
			os.Exit(1)
		}
		return
	}

	// Otherwise, run the normal CLI.
	cli.Execute()
}
