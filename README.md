# get-kube

`get-kube` (or `gkb`) is a powerful and flexible CLI tool designed to manage Kubernetes authentication tokens and configurations. With support for multiple environments—including Amazon EKS, EC2-hosted Kubernetes, and generic Kubernetes clusters—`get-kube` simplifies the process of retrieving tokens and syncing your local `.kube/config` file with remote clusters.

Whether you’re working in a cloud environment, managing on-premise infrastructure, or switching between clusters, `get-kube` has you covered.

---

## Table of Contents

- [get-kube](#get-kube)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Prerequisites](#prerequisites)
    - [Dependencies](#dependencies)
    - [Setup](#setup)
  - [Usage](#usage)
    - [Retrieve Tokens (`get`)](#retrieve-tokens-get)
      - [**EKS**](#eks)
      - [**EC2**](#ec2)
      - [**Cluster**](#cluster)
    - [Sync Kubeconfig (`sync`)](#sync-kubeconfig-sync)
      - [**Sync All Contexts**](#sync-all-contexts)
      - [**Sync a Specific Context**](#sync-a-specific-context)
  - [Building and Installation](#building-and-installation)
    - [Linux](#linux)
    - [macOS](#macos)
  - [Extending the Code](#extending-the-code)
  - [Contributing](#contributing)
  - [License](#license)

---

## Features

`get-kube` offers the following capabilities:

- **Retrieve Kubernetes authentication tokens from**:
  - **EKS**: Fetch tokens for all Amazon EKS clusters available in your AWS account.
  - **EC2**: Obtain tokens from Kubernetes clusters hosted on Amazon EC2.
  - **Generic Kubernetes Clusters**: Use SSH to retrieve tokens from any Kubernetes cluster.
- **Sync your local `.kube/config` with remote configurations**:
  - Update or merge cluster details, authentication info, and contexts into your existing kubeconfig.
  - Automatically set the current context if updated.
- **Alias support**: Use `gkb` as a shorthand for `get-kube`.
- **Lightweight and fast**: Designed with simplicity and performance in mind.

---

## Prerequisites

### Dependencies

Before using `get-kube`, ensure the following tools are installed:

- **AWS CLI**: For Amazon EKS and EC2 authentication.
- **Kubectl**: Kubernetes command-line tool.
- **SSH**: For accessing remote clusters via SSH.

You can install these dependencies using package managers (`apt`, `brew`, etc.) or by downloading them from their respective official websites.

---

### Setup

1. **AWS Configuration**: Ensure you are authenticated with your AWS account.
2. **SSH Access**: Verify that you have an SSH private key to access your EC2 or other Kubernetes hosts.
3. **Environment Variables**: Set up the following for AWS authentication:
   - `AWS_ACCESS_KEY_ID`: Your AWS access key.
   - `AWS_SECRET_ACCESS_KEY`: Your AWS secret key.
   - `AWS_DEFAULT_REGION`: Your AWS region (e.g., `us-west-1`).

---

## Usage

`get-kube` offers two primary commands: `get` and `sync`.

---

### Retrieve Tokens (`get`)

The `get` command fetches authentication tokens or kubeconfig information from EKS, EC2, or a generic Kubernetes cluster.

#### **EKS**

Fetch tokens for all available EKS clusters:

```bash
get-kube get eks
```

#### **EC2**

Fetch a token from an EC2-hosted Kubernetes cluster:

```bash
get-kube get ec2 --host <master-node-IP> --user <ssh-user>
```

#### **Cluster**

Fetch a token from a generic Kubernetes cluster via SSH:

```bash
get-kube get cluster --host <hostname> --user <ssh-user> --key <path-to-ssh-key>
```

---

### Sync Kubeconfig (`sync`)

The `sync` command synchronizes your local kubeconfig file with remote cluster data.

#### **Sync All Contexts**

Synchronize all available contexts in your local kubeconfig with remote clusters:

```bash
get-kube sync --all
```

#### **Sync a Specific Context**

Synchronize a specific context:

```bash
get-kube sync <context-name> --host <hostname> --user <ssh-user> --key <path-to-ssh-key>
```

---

## Building and Installation

To build and install `get-kube`:

### Linux

1. Build the binary:

    ```bash
    make dependencies
    make build
    ```

2. Make it executable:

    ```bash
    chmod a+x ./dist/get-kube_linux_amd64/get-kube
    ```

3. Install the binary system-wide:

    ```bash
    sudo mv ./dist/get-kube_linux_amd64/get-kube /usr/local/bin/
    ```

### macOS

1. Build the binary:

    ```bash
    make dependencies
    make build
    ```

2. Make it executable:

    ```bash
    chmod a+x ./dist/get-kube_darwin_arm64/get-kube
    ```

3. Install the binary system-wide:

    ```bash
    sudo mv ./dist/get-kube_darwin_arm64/get-kube /usr/local/bin/
    ```

---

## Extending the Code

`get-kube` is designed with extensibility in mind. To add features or modify existing functionality:

1. **Set up your development environment**:
   - Install Go (1.20 or later).
   - Clone the repository:

     ```bash
     git clone https://github.com/hemzaz/get-kube.git
     cd get-kube
     ```

2. **Modify the codebase**:
   - Add or modify commands in the `cmd/get-kube/` directory.
   - Extend functionality in the `pkg/` directory.

3. **Contribute**:
   - Open a pull request with your changes to the repository.

---

## Contributing

We welcome contributions! Please follow these steps:

1. Fork the repository.
2. Create a feature branch:

    ```bash
    git checkout -b feature/your-feature
    ```

3. Commit your changes:

    ```bash
    git commit -m "Add your feature"
    ```

4. Push to your fork:

    ```bash
    git push origin feature/your-feature
    ```

5. Open a pull request.

---

## License

`get-kube` is open-source software licensed under the MIT License. Feel free to use, modify, and distribute the code.

Authored by: hemzaz the frogodile 🐸🐊
