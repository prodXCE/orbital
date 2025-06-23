package runtime

import (
	"fmt"
	"runtime"

	"github.com/prodXCE/orbital/backends"
)

func GetManager() (ContainerManager, error) {
	switch runtime.GOOS {
	case "linux":
		// On Linux, we now return our real LinuxBackend!
		fmt.Println("[INFO] Linux detected. Using the native Linux backend.")
		return backends.NewLinuxBackend(), nil
	default:
		fmt.Printf("[INFO] OS '%s' is not yet supported. Using fallback.\n", runtime.GOOS)
		return backends.NewUnsupportedBackend(), nil
	}
}
