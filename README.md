# Go_K8_Automate

This project automates Kubernetes cluster setup using Go.

## Folder Structure

```text
k8s-cluster-creator/
├── cmd/
│   └── clusterctl/
│       └── main.go
├── internal/
│   ├── app/
│   ├── config/
│   ├── executor/
│   ├── steps/
│   ├── models/
│   ├── templates/
│   └── utils/
├── pkg/
├── deployments/
├── scripts/
├── test/
├── docs/
├── configs/
├── Makefile
├── go.mod
├── go.sum
└── README.md
```

## What Each Folder Means

### `cmd/`
Contains the application entry point.  
`cmd/clusterctl/main.go` is where the program starts when you run the tool.

### `internal/`
Contains the core logic of the project.  
Code here is private to this repository and should not be imported by external projects.

#### `internal/app/`
Holds the main application flow, orchestration, and step execution order.

#### `internal/config/`
Stores configuration loading, defaults, and validation logic.

#### `internal/executor/`
Contains code that runs commands locally or remotely on machines.

#### `internal/phases/`
Contains the automation steps for cluster setup, such as:
- update OS
- disable swap
- install container runtime
- configure container runtime
- install Kubernetes components
- initialize the cluster
- configure kubectl access
- install pod network

Each numbered folder represents one stage of the Kubernetes setup process.

#### `internal/models/`
Defines shared data structures used across the app, such as cluster details, nodes, and step results.

#### `internal/templates/`
Stores template files for generated configs such as containerd, kubeadm, sysctl, and network manifests.

#### `internal/utils/`
Contains helper functions like logging, retry logic, file handling, command execution, and error formatting.

### `pkg/`
Reserved for packages that could be reused outside the project if needed.

### `deployments/`
Holds deployment-related files and sample manifests.

### `scripts/`
Contains utility scripts for development, testing, or automation tasks.

### `test/`
Contains integration tests, end-to-end tests, and test fixtures.

### `docs/`
Stores documentation, architecture notes, and usage examples.

### `configs/`
Contains YAML or JSON configuration files for cluster setup, nodes, and runtime settings.

## Important Files

### `Makefile`
Defines common commands like build, test, lint, and run.

### `go.mod`
Defines the Go module name and dependency versions.

### `go.sum`
Locks exact dependency checksums for reproducible builds.

## Purpose of the Project

The goal is to automate Kubernetes cluster creation in a clean, repeatable, and maintainable way using Go.


k8s-cluster-creator/
├── cmd/
│   └── clusterctl/
│       └── main.go
│
├── internal/
│   ├── app/
│   │   ├── orchestrator/
│   │   ├── planner/
│   │   └── workflow/
│   │
│   ├── config/
│   │   ├── config.go
│   │   ├── defaults.go
│   │   └── validate.go
│   │
│   ├── executor/
│   │   ├── local/
│   │   ├── remote/
│   │   └── common/
│   │
│   ├── workflows/
│   │   ├── 01-update-os/
│   │   │   ├── step.go
│   │   │   ├── ubuntu.go
│   │   │   └── checks.go
│   │   │
│   │   ├── 02-disable-swap/
│   │   │   ├── step.go
│   │   │   ├── swapoff.go
│   │   │   └── checks.go
│   │   │
│   │   ├── 03-install-container-runtime/
│   │   │   ├── step.go
│   │   │   ├── containerd.go
│   │   │   └── checks.go
│   │   │
│   │   ├── 04-configure-containers/
│   │   │   ├── step.go
│   │   │   ├── sysctl.go
│   │   │   ├── kernel_modules.go
│   │   │   ├── cgroup.go
│   │   │   └── checks.go
│   │   │
│   │   ├── 05-install-k8s-components/
│   │   │   ├── step.go
│   │   │   ├── kubeadm.go
│   │   │   ├── kubelet.go
│   │   │   ├── kubectl.go
│   │   │   └── checks.go
│   │   │
│   │   ├── 06-initialize-cluster/
│   │   │   ├── step.go
│   │   │   ├── init_control_plane.go
│   │   │   ├── token.go
│   │   │   └── checks.go
│   │   │
│   │   ├── 07-configure-kubectl-access/
│   │   │   ├── step.go
│   │   │   ├── kubeconfig.go
│   │   │   └── checks.go
│   │   │
│   │   └── 08-install-pod-network/
│   │       ├── step.go
│   │       ├── calico.go
│   │       ├── cilium.go
│   │       └── checks.go
│   │
│   ├── models/
│   │   ├── cluster.go
│   │   ├── node.go
│   │   ├── step.go
│   │   └── result.go
│   │
│   ├── templates/
│   │   ├── containerd/
│   │   ├── kubeadm/
│   │   ├── sysctl/
│   │   └── networking/
│   │
│   └── utils/
│       ├── logger/
│       ├── retry/
│       ├── shell/
│       ├── files/
│       └── errors/
│
├── pkg/
│   ├── api/
│   └── version/
│
├── deployments/
│   ├── sample-configs/
│   └── manifests/
│
├── scripts/
│   ├── dev/
│   └── test/
│
├── test/
│   ├── integration/
│   ├── e2e/
│   └── fixtures/
│
├── docs/
│   ├── architecture/
│   ├── workflow/
│   └── examples/
│
├── configs/
│   ├── cluster.yaml
│   ├── nodes.yaml
│   └── runtime.yaml
│
├── Makefile
├── go.mod
├── go.sum
└── README.md

