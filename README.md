# Orbital - A Portable Container Engine in Go

Orbital is a learning project to build a lightweight, cross-platform containerization tool from scratch in Go. The goal is to demystify the core concepts of process isolation (like namespaces, jails, and job objects) by creating a functional, portable container engine.

This project is built with a modular, pluggable architecture, allowing different OS-specific isolation backends to be used by a single, unified runtime engine.

## Current Status: Phase 3 - Linux Backend with Namespaces

This phase marks a major milestone: **Orbital can now run its first truly isolated container on Linux.** We have built our first functional backend using Linux's native process isolation capabilities.

The core architectural work from Phase 2 has paid off, allowing us to "plug in" this new Linux backend without changing the CLI or the overall program flow.

## Features Implemented

- **Functional Linux Backend:** A native backend that uses Go's `syscall` package to interact directly with the Linux kernel.
- **Process & Filesystem Isolation:**
  - **Namespaces:** The backend creates new `UTS` (for hostname), `PID` (for process IDs), and `Mount` namespaces for the container.
  - `chroot`: The container's root filesystem is changed to a dedicated directory, preventing it from accessing the host's filesystem.
- **Go Build Tags:** The Linux-specific code is isolated using the `//go:build linux` tag, ensuring the project remains buildable on other operating systems like Windows and macOS.
- **Parent/Child Process Execution:** Implemented the standard container pattern where the main process sets up the namespaces and a "child" process executes within them.

## How to Build and Run (Linux)

**Note:** Running a container currently requires a Linux host, `sudo` privileges, and Docker (to create the initial root filesystem).

### 1. Prerequisites (Linux Host)

- Go (version 1.18 or later)
- Git
- Docker

### 2. One-Time Root Filesystem Setup

The container needs its own root filesystem (`/`). We can create a minimal one using an Alpine Linux image from Docker.

**Run this in your terminal:**

```bash
# Create a directory for the rootfs
mkdir -p /tmp/orbital-rootfs

# Use Docker to export the Alpine filesystem into our new directory
docker export $(docker create alpine) | sudo tar -C /tmp/orbital-rootfs -xvf -
```

### 3. Build the Orbital Binary

```bash
# Clone the repository if you haven't already
git clone https://github.com/prodXCE/orbital.git
cd orbital

# Tidy dependencies and build the executable
go mod tidy
go build -o orbital
```

### 4. Run Your First Container!

You must use `sudo` because creating namespaces and using `chroot` are privileged operations.

```bash
sudo ./orbital run /bin/sh
```

## Expected Output & Verification

Upon running the command, you will be dropped into a new shell prompt, which is running inside your isolated container:

```
[INFO] Linux detected. Using the native Linux backend.
--> Using backend of type: *backends.LinuxBackend
[Linux Backend] Starting container for command: /bin/sh []
--> Entering child process execution mode...
[Child Process] Running inside container! Command: /bin/sh []
Container started with PID: 31337
/ #
```

You are now inside the container. You can verify the isolation:

- **Check the hostname:**

```bash
/ # hostname
orbital-container
```

- **Check the running processes:**

```bash
/ # ps aux
```

You will see that your shell (`/bin/sh`) has PID 1. You cannot see any processes from the host OS.

- **Check the root directory:**

```bash
/ # ls /
bin    dev    etc    home   lib    media  mnt    proc   root   ...
```

This lists the contents of `/tmp/orbital-rootfs`, not your host's `/`.

To exit the container and return to your host shell, simply type `exit`.
