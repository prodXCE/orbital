package runtime

// Container is a data structure representing a container.
// It is now separate from the interface definition that backends use,
// which breaks the import cycle. The CLI or runtime can use this
// struct to hold container information.
type Container struct {
	ID string
}

// ContainerManager defines the set of operations for managing a container's lifecycle.
type ContainerManager interface {
	// Start creates and starts a new container for the given command.
	// IT NOW RETURNS a container ID (string) and an error. This is the key change
	// that breaks the import cycle, as backends no longer need to know about
	// the 'Container' struct.
	Start(command string, args []string) (string, error)

	Stop(containerID string) error
	Status(containerID string) (string, error)
}
