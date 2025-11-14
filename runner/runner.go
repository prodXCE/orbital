/*
 * This package 'runner' contains all the "parent" logic
 */

package runner

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"syscall"
)

func Run(rootfsPath string, hostname string, args []string) {
	fmt.Printf("Parent: Running command %v in %s with hostname %s\n", args, rootfsPath, hostname)

	absRootfs, err := filepath.Abs(rootfsPath)
	if err != nil {
		fmt.Printf("Parent: Error resolving rootfs path: %v\n", err)
		os.Exit(1)
	}

	exePath, err := os.Executable()
	if err != nil {
		fmt.Printf("Parent: Error finding executable; %v\n", err)
		os.Exit(1)
	}

	childArgs := append([]string{"child", absRootfs, hostname}, args...)

	cmd := exec.Command(exePath, childArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWNS |
			syscall.CLONE_NEWPID |
			syscall.CLONE_NEWUTS,
	}

	if err := cmd.Run(); err != nil {
		fmt.Printf("Parent: Error running child process: %v\n", err)
		os.Exit(1)
	}

}
