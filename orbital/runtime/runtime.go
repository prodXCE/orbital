package runtime

import (
	"fmt"
	"runtime"

	"github.com/prodXCE/orbital/backends"
)

func GetManager() (ContainerManager, error) {
	switch runtime.GOOS {
	case "linux":
		fmt.Println("[INFO] Linux detected. Using the native Linux backend.")
		return backends.NewLinuxBackend(), nil
	case "windows":
		// On Windows, we now return our real WindowsBackend!
		fmt.Println("[INFO] Windows detected. Using the native Windows backend.")
		return backends.NewWindowsBackend(), nil
	default:
		fmt.Printf("[INFO] OS '%s' is not yet supported. Using fallback.\n", runtime.GOOS)
		return backends.NewUnsupportedBackend(), nil
	}
}