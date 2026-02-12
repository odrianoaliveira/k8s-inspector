# k8s-inspector

A lightweight Kubernetes cluster inspection CLI tool built with Go.

## Overview

k8s-inspector is a command-line utility designed to help developers and operators inspect and monitor Kubernetes clusters. It provides quick access to pod information, status, resource usage, and logs without the overhead of complex dashboards.

## Project Plan

### Core Features

- **List all pods** in a namespace or across all namespaces
- **Show pod status** (Running, Pending, CrashLoopBackOff, etc.)
- **Display resource usage** (CPU/memory requests/limits)
- **Export pod logs** to a file for debugging

## Tech Stack

### Go Libraries

- **client-go**: Official Kubernetes client library for Go
  - Provides access to Kubernetes API resources
  - Handles authentication and cluster communication

- **cobra**: CLI framework for building command-line applications
  - Used by kubectl, helm, and other popular Kubernetes tools
  - Provides subcommands, flags, and argument parsing

## Getting Started

### Prerequisites

- Go 1.26 or higher
- Access to a Kubernetes cluster (kubeconfig configured)

### Installation

Clone the repository:
```bash
git clone https://github.com/yourusername/k8s-inspector.git
cd k8s-inspector
```

Build the application:
```bash
make build
```

### Usage

#### List Pods

View all pods in a specific namespace:
```bash
./k8s-inspector pods --namespace default
```

View all pods across all namespaces:
```bash
./k8s-inspector pods --all-namespaces
```

#### Available Commands

Run `./k8s-inspector --help` to see all available commands and options.

## Development

### Build the Project

```bash
make build
```

## Project Structure

```
k8s-inspector/
├── main.go           # Entry point
├── cmd/
│   ├── root.go       # Root command definition
│   └── pods.go       # Pods command implementation
├── go.mod            # Go module definition
├── go.sum            # Go module checksums
├── Makefile          # Build automation
└── README.md         # This file
```

## Roadmap

### Phase 1 (Current)
- [x] Project structure
- [ ] List pods in namespace
- [ ] Display pod status
- [ ] Show resource usage

### Phase 2
- [ ] Export pod logs to file
- [ ] Advanced filtering options
- [ ] Resource metrics display

### Phase 3
- [ ] Node inspection
- [ ] Persistent volume inspection
- [ ] Service discovery
- [ ] Configuration validation
