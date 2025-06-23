Orbital - A Portable Container Engine in Go

Orbital is a learning project to build a lightweight, cross-platform containerization tool from scratch in Go. The goal is to demystify the core concepts of process isolation (like namespaces, jails, and job objects) by creating a functional, portable container engine.

This project is built with a modular, pluggable architecture, allowing different OS-specific isolation backends to be used by a single, unified runtime engine.
Current Status: Phase 2 - Core Architecture & Interface

This phase implements the core architectural foundation of Orbital. The project now has a clean, pluggable system ready to accept OS-specific backends. We have defined a universal ContainerManager interface that all backends must adhere to and created a "factory" that selects the appropriate backend at runtime.

No actual containerization occurs yet, but running the run command now demonstrates that the architectural model is working correctly by selecting a fallback "unsupported" backend and exiting gracefully.
Features Implemented

    ContainerManager Interface: A clear Go interface defines the required behavior for all backends (Start, Stop, Status).

    Pluggable Backend Factory: A factory function (runtime.GetManager()) automatically detects the host OS and selects the correct backend.

    Fallback "Unsupported" Backend: A safe, default backend is used for any operating system that doesn't have a specific implementation yet, ensuring the program behaves predictably.

    Import Cycle Resolution: The architecture has been refined to prevent circular dependencies between packages, which is a key principle of good Go design.

How to Build and Run
Prerequisites

    Go (version 1.18 or later)

    Git

Build Steps

    Clone the repository (if you haven't already):

    git clone https://github.com/prodXCE/orbital.git
    cd orbital

    Fetch Go dependencies:

    go mod tidy

    Build the binary:
    This command compiles the project and creates an executable named orbital in the root directory.

    go build -o orbital

Usage Examples

Running the tool now demonstrates the new architecture. The program correctly identifies that no real backend is available and prints an error, proving the system works as designed.

1. Attempt to run a command on Linux:

./orbital run /bin/sh

Expected Output:

[INFO] Linux detected. The real backend will be implemented in Phase 3.
--> Using backend of type: *backends.UnsupportedBackend
Error starting container: container operations are not supported on this OS: linux

2. Attempt to run a command on macOS or Windows:

./orbital run cmd.exe

Expected Output (on Windows):

[INFO] OS 'windows' is not yet supported. Using fallback.
--> Using backend of type: *backends.UnsupportedBackend
Error starting container: container operations are not supported on this OS: windows

