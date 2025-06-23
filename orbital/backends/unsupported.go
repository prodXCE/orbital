package backends

import (
	"errors"
	"fmt"
	"runtime" // This now only imports the standard library 'runtime'.
)

// UnsupportedBackend is a struct that implements the ContainerManager interface.
type UnsupportedBackend struct{}

// NewUnsupportedBackend creates a new instance of the UnsupportedBackend.
func NewUnsupportedBackend() *UnsupportedBackend {
	return &UnsupportedBackend{}
}

// Start now matches the updated interface, returning (string, error).
// It no longer needs to import our project's runtime package at all.
func (b *UnsupportedBackend) Start(command string, args []string) (string, error) {
	errMsg := fmt.Sprintf("container operations are not supported on this OS: %s", runtime.GOOS)
	// Return an empty string for the ID and the error.
	return "", errors.New(errMsg)
}

// Stop returns an error.
func (b *UnsupportedBackend) Stop(containerID string) error {
	errMsg := fmt.Sprintf("Stop operation is not supported on this OS: %s", runtime.GOOS)
	return errors.New(errMsg)
}

// Status returns an error.
func (b *UnsupportedBackend) Status(containerID string) (string, error) {
	errMsg := fmt.Sprintf("Status operation is not supported on this OS: %s", runtime.GOOS)
	return "", errors.New(errMsg)
}

