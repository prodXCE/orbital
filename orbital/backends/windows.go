package backends

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"unsafe"

	"golang.org/x/sys/windows"
)

// WindowsBackend is the struct that implements the ContainerManager interface for Windows.
type WindowsBackend struct{}

// NewWindowsBackend is the factory function for our Windows backend.
func NewWindowsBackend() *WindowsBackend {
	return &WindowsBackend{}
}

// Start implements the container start operation for Windows using Job Objects.
func (b *WindowsBackend) Start(command string, args []string) (string, error) {
	fmt.Printf("[Windows Backend] Starting container for command: %s %v\n", command, args)

	// 1. Create a new Job Object.
	// This object will hold our containerized process and any children it spawns.
	job, err := windows.CreateJobObject(nil, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create job object: %w", err)
	}
	defer windows.CloseHandle(job)

	// 2. Configure the Job Object.
	// We'll set a basic limit and configure it to terminate all processes when the handle is closed.
	info := &windows.JOBOBJECT_EXTENDED_LIMIT_INFORMATION{
		BasicLimitInformation: windows.JOBOBJECT_BASIC_LIMIT_INFORMATION{
			LimitFlags: windows.JOB_OBJECT_LIMIT_KILL_ON_JOB_CLOSE,
		},
	}
	if _, err := windows.SetInformationJobObject(job, windows.JobObjectExtendedLimitInformation, uintptr(unsafe.Pointer(info)), uint32(unsafe.Sizeof(*info))); err != nil {
		return "", fmt.Errorf("failed to set job object information: %w", err)
	}

	// 3. Create the command to execute.
	cmd := exec.Command(command, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// This is important: We need to create the process with a CREATE_SUSPENDED flag
	// so that we have a chance to assign it to our job object before it starts running.
	cmd.SysProcAttr = &windows.SysProcAttr{
		CreationFlags: windows.CREATE_SUSPENDED,
	}

	// 4. Start the command (in a suspended state).
	if err := cmd.Start(); err != nil {
		return "", fmt.Errorf("failed to start suspended process: %w", err)
	}

	// 5. Assign the newly created process to our Job Object.
	// We get the process handle from the running command.
	processHandle := windows.Handle(cmd.Process.Handle())
	if err := windows.AssignProcessToJobObject(job, processHandle); err != nil {
		return "", fmt.Errorf("failed to assign process to job object: %w", err)
	}

	// 6. Now that the process is safely inside the job object, we can resume its execution.
	p, _ := cmd.Process.Find()
	if _, err := windows.ResumeThread(windows.Handle(p.Handle)); err != nil {
		return "", fmt.Errorf("failed to resume process thread: %w", err)
	}

	containerID := strconv.Itoa(cmd.Process.Pid)
	fmt.Printf("Container started with PID: %s and assigned to Job Object\n", containerID)

	// Wait for the command to finish.
	if err := cmd.Wait(); err != nil {
		fmt.Printf("Container process exited with error: %v\n", err)
	}
	
	return containerID, nil
}


// Stop is a placeholder for now.
func (b *WindowsBackend) Stop(containerID string) error {
	return fmt.Errorf("stop is not yet implemented")
}

// Status is a placeholder for now.
func (b *WindowsBackend) Status(containerID string) (string, error) {
	return "", fmt.Errorf("status is not yet implemented")
}
