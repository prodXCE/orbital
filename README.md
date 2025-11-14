# orbital 

*A minimal container runtime written from scratch in Go.*

**orbital** is a learning-focused container runtime that demonstrates the *actual* Linux kernel mechanisms behind Docker, Podman, and other runtimes.

It is **not** meant to replace existing tools. Instead, it exists to teach what a container really is: **just an isolated Linux process sharing the host kernel**.

---

## 🏛️ Core Idea: What *Is* a Container?

A container is **not** a virtual machine.
There is no virtual hardware, no hypervisor, no separate kernel.

A container is simply:

> **A regular Linux process running with restricted visibility and access
> using namespaces, chroot, and filesystem mounts.**

orbital demonstrates these Linux kernel features:

### 🔹 Linux Namespaces

Namespaces isolate resources such that a process believes it owns them:

* **Mount Namespace (CLONE_NEWNS)**
  Gives the container its own independent mount table.

* **PID Namespace (CLONE_NEWPID)**
  The first process inside the namespace becomes **PID 1**.
  It cannot see host PIDs.

* **UTS Namespace (CLONE_NEWUTS)**
  Allows a custom hostname inside the container.

### 🔹 Filesystem Isolation: `chroot`

`chroot` changes the process’s root directory (`/`) to a new filesystem, preventing escape into the host.

### 🔹 Go syscalls

orbital uses Go’s `syscall` and `unix` packages as a thin wrapper over:

* `clone()`
* `mount()`
* `chroot()`
* `exec()`
* namespace flags + mount options

---

## 🧠 How orbital Works — The Parent/Child Re-exec Pattern

The core of orbital is the **re-exec trick** used by Docker, runc, containerd, and others.

### 1. **Parent Process** (the normal CLI)

When you run:

```bash
./orbital run alpine-arm /bin/sh
```

The parent process:

1. Parses CLI flags.
2. Prepares clone flags (e.g., `CLONE_NEWPID | CLONE_NEWNS | CLONE_NEWUTS`).
3. Uses `exec.Command()` to **re-execute itself** with a magic argument: `child`.
   Example:

   ```
   /opt/orbital/orbital child <args...>
   ```

### 2. **Gatekeeper (`main.go`)**

A fresh orbital process starts.

`main.go` checks:

* If it sees `child` → it knows it must run container isolation code.
* Otherwise → it runs the normal CLI handlers.

### 3. **Child Process** (inside namespaces)

Inside new namespaces, the child:

1. Makes mount propagation private (`MS_PRIVATE`).
2. Performs the `chdir()` + `chroot()` jail setup.
3. Mounts needed filesystems:

   * `/proc`
   * `/tmp`
4. Sets the container’s hostname.
5. Finally calls `syscall.Exec()` to replace itself with the user’s command.

At this moment, the Go process disappears — you're left with a real isolated process that thinks it's running on its own machine.

---

## ✨ Features

* Run processes inside **Mount**, **PID**, and **UTS** namespaces
* Custom hostnames using `--hostname` / `-H`
* Root filesystem downloader:

  * `orbital pull <image>` downloads and extracts rootfs images
* Smart image selection:

  * `orbital run alpine-arm /bin/sh`
  * `orbital run ./myrootfs /bin/sh`
* Clean, modular architecture:

  * `cmd/` — CLI (Cobra)
  * `runner/` — parent-side namespace creation
  * `isolation/` — child-side jail setup
  * `downloader/` — rootfs download and extraction

---

## 🚀 Getting Started

### 1. Prerequisites

* **Go 1.18+**
* **Linux only**
  (Uses syscalls not available on macOS/Windows.)
* **Root access (sudo)**
  Needed for `clone()`, `mount()`, `chroot()`.
* **Correct architecture rootfs**

  * ARM64 rootfs for ARM hosts
  * AMD64 rootfs for Intel/AMD hosts

### 2. Install orbital

Clone the repository:

```bash
git clone https://github.com/prodXCE/orbital.git
cd orbital
```

Build the binary:

```bash
go build -o orbital .
```

If your home directory is mounted with `noexec`, move the binary:

```bash
sudo mv orbital /opt/orbital/orbital
```

Ignore pulled images:

```bash
echo ".orbital/" >> .gitignore
```

---

## 🧰 Usage

### Step 1: Pull an Image

This downloads a rootfs into `.orbital/images/...`.

```bash
./orbital pull alpine-arm
```

For x86_64:

```bash
./orbital pull alpine-amd
```

### Step 2: Run a Container (requires sudo)

#### Interactive shell

```bash
sudo ./orbital run alpine-arm /bin/sh
```

#### Custom hostname

```bash
sudo ./orbital run --hostname my-test-box alpine-arm ps aux
```

#### Using a local rootfs path

```bash
sudo ./orbital run ./my-rootfs /bin/sh
```

### Step 3: Help

```bash
./orbital --help
./orbital run --help
./orbital pull --help
```

---

## 📁 Project Structure

```
orbital/
├── .gitignore
├── .orbital/             # Pulled images (ignored from Git)
│   └── images/
│       └── alpine-arm/
├── go.mod
├── go.sum
├── main.go               # Gatekeeper (detects 'child')
│
├── cmd/                  # CLI using Cobra
│   ├── root.go
│   ├── run.go
│   └── pull.go
│
├── downloader/           # Image pull + extract
│   └── downloader.go
│
├── isolation/            # Child process (inside namespaces)
│   └── isolation.go
│
└── runner/               # Parent logic (create namespaces)
    └── runner.go
```
