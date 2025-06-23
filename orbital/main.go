package main

import (
	"os"
	"runtime"

	"github.com/prodXCE/orbital/backends"
	"github.com/prodXCE/orbital/cli"
)

func main() {
	// The child process re-execution is a Linux-specific pattern for our design.
	// This logic will not run on Windows, which is correct.
	if len(os.Args) > 1 && os.Args[1] == "child" {
		if runtime.GOOS == "linux" {
			backends.ChildProcess(os.Args[2], os.Args[3:])
			return
		}
	}

	// All other OSes will run the CLI directly.
	cli.Execute()
}
