# get-kube

`get-kube` is a CLI tool designed to retrieve Kubernetes authentication tokens from various environments and update the local `.kube/config` file. It supports EKS clusters, EC2-hosted Kubernetes, Kind clusters, and Docker Desktop Kubernetes.

## Table of Contents

- [Prerequisites](#prerequisites)
  - [Dependencies](#dependencies)
  - [Setup](#setup)
- [Usage](#usage)
  - [EKS](#eks)
  - [EC2](#ec2)
  - [Kind](#kind)
  - [Docker Desktop](#docker-desktop)
- [Building and Installation](#building-and-installation)
  - [Linux](#linux)
  - [macOS](#macos)
- [Extending the Code](#extending-the-code)

## Prerequisites

### Dependencies

Before using `get-kube`, ensure you have the following tools installed:

- Docker: Used for containerized environments.
- Kind: For running local Kubernetes clusters using Docker container nodes.
- AWS CLI: Required for interacting with Amazon EKS and EC2.
- Kubectl: Kubernetes command-line tool.

You can install these dependencies using package managers like `apt` for Linux, `brew` for macOS, or by downloading them from their respective official websites.

### Setup

- **AWS Account**: Ensure you have an AWS account and are authenticated. This is necessary for fetching tokens from EKS and EC2-hosted Kubernetes clusters.
  
- **Environment Variables**: Set up the following environment variables:
  - `AWS_ACCESS_KEY_ID`: Your AWS access key.
  - `AWS_SECRET_ACCESS_KEY`: Your AWS secret key.
  - `AWS_DEFAULT_REGION`: The default AWS region (e.g., `us-west-1`).

## Usage

### EKS

To fetch tokens from all EKS clusters you have access to:

```bash
get-kube eks
```

### EC2

To fetch tokens from a Kubernetes cluster hosted on an EC2 instance, provide the master node's name, IP, or FQDN:

```bash
get-kube ec2 <master-node-name/IP/FQDN>
```

### Kind

To fetch tokens from all local Kind clusters:

```bash
get-kube kind
```

### Docker Desktop

To fetch tokens from Docker Desktop Kubernetes:

```bash
get-kube docker-desktop
```

Optionally, specify a context name:

```bash
get-kube docker-desktop --context=my-context-name
```

## Building and Installation

### Linux

To build for Linux:

```bash
make build
```

After building, make the binary executable:

```bash
chmod a+x ./dist/get-kube_linux_amd64/get-kube
```

Then, install the binary:

```bash
sudo mv ./dist/get-kube_linux_amd64/get-kube /usr/local/bin/
```

### macOS

To build for macOS:

```bash
make build
```

After building, make the binary executable:

```bash
chmod a+x ./dist/get-kube_darwin_arm64/get-kube
```

Then, install the binary:

```bash
sudo mv ./dist/get-kube_darwin_arm64/get-kube /usr/local/bin/
```

# Extending the Code

If you wish to extend the functionality of `get-kube`, follow these steps:

1. **Setup the Development Environment**: Ensure you have Go installed and set up on your machine.
2. **Clone the Repository**: `git clone https://github.com/hemzaz/get-kube.git`
3. **Navigate to the Code Directory**: `cd get-kube`
4. **Make Changes**: Add new features or modify existing ones in the `pkg` directory or the main application in the `cmd` directory.
5. **Build and Test**: Use the provided `Makefile` to build and test your changes.
6. **Contribute**: If you believe your changes could benefit others, consider creating a pull request to merge your changes back into the main repository.

---

Authored by: **hemzaz the frogodile** 🐸🐊
#
