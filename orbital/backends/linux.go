package backends

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
)

// The path to the root filesystem for our container.
// In a real tool, this would be configurable.
const rootfsPath = "/tmp/orbital-rootfs"

// LinuxBackend is the struct that implements the ContainerManager interface for Linux.
type LinuxBackend struct{}

// NewLinuxBackend is the factory function for our Linux backend.
func NewLinuxBackend() *LinuxBackend {
	return &LinuxBackend{}
}

func IsLinux() bool {
    return runtime.GOOS == "linux"
}

// Start implements the container start operation for Linux.
func (b *LinuxBackend) Start(command string, args []string) (string, error) {
	fmt.Printf("[Linux Backend] Starting container for command: %s %v\n", command, args)

	// Create a new exec.Cmd to run the user's command.
	// We are actually running our own program again, but with a special argument 'child'.
	// This is a common pattern for setting up namespaces.
	cmd := exec.Command("/proc/self/exe", append([]string{"child", command}, args...)...)

	// Configure the new process to run in new namespaces.
	// This is where the core isolation happens.
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | // New UTS namespace for hostname
			syscall.CLONE_NEWPID | // New PID namespace for processes
			syscall.CLONE_NEWNS, // New mount namespace for filesystem
	}

	// Connect the standard I/O of the container to our terminal.
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// Start the command. This will create the new process in the new namespaces.
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start container process: %w", err)
	}

	// For now, we'll just use the process ID as the container ID.
	containerID := fmt.Sprintf("%d", cmd.Process.Pid)
	fmt.Printf("Container started with PID: %s\n", containerID)

	// Wait for the container process to exit.
	// In a real tool, we would manage this process in the background.
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Container process exited with error: %v\n", err)
	}

	return containerID, nil
}

// childProcess is the function that runs *inside* the container's namespaces.
// It sets up the new root filesystem and executes the user's command.
func ChildProcess(command string, args []string) {
	fmt.Printf("[Child Process] Running inside container! Command: %s %v\n", command, args)

	// 1. Set a new hostname for the container.
	must(syscall.Sethostname([]byte("orbital-container")))

	// 2. Change the root filesystem for the process.
	must(syscall.Chroot(rootfsPath))
	must(os.Chdir("/")) // Change directory to the new root.

	// 3. Mount necessary filesystems like /proc.
	// /proc is a virtual filesystem that provides process and kernel information.
	// It's essential for many tools like 'ps' or 'top' to work correctly.
	must(syscall.Mount("proc", "proc", "proc", 0, ""))
	defer must(syscall.Unmount("proc", 0)) // Unmount when done.

	// 4. Finally, execute the user's command.
	// syscall.Exec replaces the current process with the new one.
	// We need to find the absolute path of the command.
	cmdPath, err := exec.LookPath(command)
	if err != nil {
		fmt.Printf("Error looking up command %s: %v\n", command, err)
		os.Exit(1)
	}

	err = syscall.Exec(cmdPath, append([]string{command}, args...), os.Environ())
	if err != nil {
		// This part will only be reached if Exec fails.
		fmt.Printf("Error executing command: %v\n", err)
		os.Exit(1)
	}
}

// A simple helper function to panic on errors for setup steps.
// This keeps the setup code cleaner. In a production tool, you'd handle these errors.
func must(err error) {
	if err != nil {
		panic(err)
	}
}

// Stop is a placeholder for now.
func (b *LinuxBackend) Stop(containerID string) error {
	return fmt.Errorf("stop is not yet implemented")
}

// Status is a placeholder for now.
func (b *LinuxBackend) Status(containerID string) (string, error) {
	return "", fmt.Errorf("status is not yet implemented")
}
