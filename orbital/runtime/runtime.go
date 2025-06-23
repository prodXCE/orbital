package runtime

import (
	"fmt"
	"runtime"

	"github.com/prodXCE/orbital/backends"
)

// GetManager is our backend factory. Its logic remains correct.
func GetManager() (ContainerManager, error) {
	switch runtime.GOOS {
	case "linux":
		fmt.Println("[INFO] Linux detected. The real backend will be implemented in Phase 3.")
		return backends.NewUnsupportedBackend(), nil
	default:
		fmt.Printf("[INFO] OS '%s' is not yet supported. Using fallback.\n", runtime.GOOS)
		return backends.NewUnsupportedBackend(), nil
	}
}

