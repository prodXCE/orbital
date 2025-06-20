# Orbital – A Portable Container Engine in Go

**Orbital** is a hands-on learning project focused on building a lightweight, cross-platform containerization tool from scratch using Go. The goal is to **demystify process isolation techniques** like namespaces, jails, and job objects by creating a functional, portable container engine.

Orbital is designed with a **modular, pluggable architecture** that can support different OS-specific isolation backends under a unified runtime engine.

---

## 🚀 Project Status: Phase 1 – CLI Foundation & Skeleton

This phase establishes the basic project structure. The application:

* Compiles successfully
* Provides a functional **Command-Line Interface (CLI)**
* Includes a **placeholder runtime engine** with OS detection

> **Note:** No actual containerization or process isolation is implemented yet. This is an architecture and CLI preparation phase.

---

## ✨ Key Features

* **Modular Project Structure**
  Clean separation of concerns into:

  * `cli` package
  * `runtime` package
  * `backends` package

* **Professional CLI with Cobra**
  Supports:

  * `orbital run [command]` → Placeholder for running commands
  * `orbital status [containerID]` → Placeholder for checking status
  * Auto-generated help text with `--help`

* **Host OS Detection**
  The runtime accurately detects the underlying operating system (Linux, Windows, macOS, etc.).

* **Go Module Support**
  Fully initialized as a Go module with proper dependency management.

---

## 🔧 Getting Started

### Prerequisites

* Go (version **1.18** or later)
* Git

### Setup & Build Instructions

1. **Clone the Repository**
   `git clone https://github.com/prodXCE/orbital.git`

2. **Navigate to Project Directory**
   `cd orbital`

3. **Fetch Dependencies**
   `go mod tidy`

4. **Build the Project**
   `go build -o orbital`

   This will generate the `orbital` executable in the root directory.

---

## 📦 Usage Examples

### 1. View CLI Help

```bash
./orbital --help
```

### 2. Run a Command (Simulated)

```bash
./orbital run /bin/sh -c "echo hello from inside the container"
```

**Expected Output:**

```
========================================
Host OS detected: linux
========================================
Starting container for command: '/bin/sh' with args: -c "echo hello from inside the container"

[INFO] Phase 1 complete. No container was actually started.
```

### 3. Check Container Status (Simulated)

```bash
./orbital status my-container-123
```

**Expected Output:**

```
Checking status for container: my-container-123...
[INFO] Status command is not implemented yet.
```

---

## 🗺️ Project Roadmap

* ✅ Phase 1: Project Skeleton & CLI Foundation
* ⏳ Phase 2: Basic Process Isolation (namespaces, minimal sandbox)
* ⏳ Phase 3: Filesystem Isolation (chroot, overlay filesystems)
* ⏳ Phase 4: Networking (virtual interfaces, port mappings)
* ⏳ Phase 5: OS-Specific Backends (platform-level integration)

---

## 🤝 Contribution Guidelines

This project is intended for educational purposes, but contributions are always welcome.
You can:

* Fork the repository
* Open issues for bugs, improvements, or feature ideas
* Submit pull requests

Your feedback will help improve the project and refine the learning process.

---

## 📄 License

This project is open-source and available under the [MIT License](LICENSE).
