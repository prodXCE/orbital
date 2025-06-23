# Orbital - A Portable Container Engine in Go

Orbital is a learning project to build a lightweight, cross-platform containerization tool from scratch in Go. The goal is to demystify the core concepts of process isolation by creating a functional, portable container engine that runs on multiple operating systems.

This project is built with a modular, pluggable architecture, allowing different OS-specific isolation backends to be used by a single, unified runtime engine.

## Current Status: Phase 4 - Cross-Platform Support (Linux & Windows)

This phase achieves a major architectural goal: **Orbital is now a true cross-platform application.** We have successfully implemented a second native backend for the Windows operating system, which runs alongside our existing Linux backend.

The `orbital` binary can be compiled on either Linux or Windows, and it will automatically select the correct native backend to provide process isolation.

## Features Implemented

### Linux Backend

- **Full Isolation:** Uses Linux Namespaces (`UTS`, `PID`, `Mount`, `Network`) for high-level isolation.
- **Filesystem Sandboxing:** Uses `chroot` to confine the container to a dedicated root filesystem.
- **Virtual Networking:** Creates a network bridge and veth pairs to provide containers with a private IP address and internet access via NAT on the host.

### Windows Backend (New!)

- **Process Grouping with Job Objects:** Uses native Windows Job Objects to group the container process and any of its children. This ensures that when the main process is terminated, all child processes are also cleaned up.
- **Process-Level Isolation:** Provides a lightweight form of isolation focused on process lifecycle management.
- **Build-Tag Separation:** Uses Go build tags (`//go:build windows`) to ensure the Windows-specific code is completely separate from the Linux code.
- **Note:** This backend does _not_ provide filesystem or network isolation. The containerized process shares the host's filesystem and network stack, which is typical for basic process containers on Windows without Hyper-V.

## How to Build and Run

### Prerequisites

- Go (version 1.18 or later)
- Git

### On Linux

**1. One-Time Setup:** You will need `docker` to create the rootfs and `bridge-utils` for networking.

```bash
# Install bridge utilities
sudo apt-get update && sudo apt-get install bridge-utils

# Create the root filesystem
mkdir -p /tmp/orbital-rootfs
docker export $(docker create alpine) | sudo tar -C /tmp/orbital-rootfs -xvf -
```

**2. Build the Binary:**

```bash
git clone https://github.com/prodXCE/orbital.git
cd orbital
go build -o orbital
```

**3. Run a Container:** `sudo` is required for creating namespaces and configuring the network.

```bash
sudo ./orbital run /bin/sh
```

This will place you in a fully isolated shell with its own IP address.

### On Windows

**1. Build the Binary:** Open a PowerShell or Command Prompt.

```powershell
git clone https://github.com/your-username/orbital.git
cd orbital
go build -o orbital.exe
```

**2. Run a Container:** No special privileges are required.

```powershell
./orbital.exe powershell
```

This will start a new PowerShell process that is managed within a Windows Job Object. You can exit by typing `exit`.
