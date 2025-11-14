package main

import (
	"fmt"
	"os"

	"github.com/prodXCE/orbital/cmd"
	"github.com/prodXCE/orbital/isolation"
)

func main() {

	if len(os.Args) > 1 && os.Args[1] == "child" {
		fmt.Println("Child: Detected, running isolation...")

		if len(os.Args) < 5 {
			fmt.Println("child: Not enough args for child process")
			os.Exit(1)
		}

		isolation.Child(os.Args[2], os.Args[3], os.Args[4:])
		return
	}

	cmd.Execute()
}
