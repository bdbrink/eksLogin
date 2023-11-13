# EKS Login Script

## Overview

This Go script provides a convenient way to authenticate with an Amazon Elastic Kubernetes Service (EKS) cluster. It generates and executes the necessary AWS CLI command to update the kubeconfig file.

## Prerequisites

- [AWS CLI](https://aws.amazon.com/cli/) installed and configured with the necessary credentials.
- [Amazon EKS CLI](https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html) for cluster management.

## Usage

```bash
go run main.go {hangar} {env} {region}
```