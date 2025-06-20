package runtime

import (
	"fmt"
	"runtime"
	"strings"
)

func RunContainer(command string, args []string) {
	fmt.Println("============================================")
	fmt.Printf("HOST OS DETECTED: %s\n", runtime.GOOS)
	fmt.Println("============================================")

	fmt.Printf("Starting container for command: '%s' with args: %s\n", command, strings.Join(args, " "))

	fmt.Println("\n[INFO] Phase 1 completed. No container started")
}
