# Rabbitmq Operator

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/from-referrer/)
![CI build and Deploy](https://github.com/kokuwaio/rabbitmq-operator/workflows/CI%20build%20and%20Deploy/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/kokuwaio/rabbitmq-operator/badge.svg?branch=master)](https://coveralls.io/github/kokuwaio/rabbitmq-operator?branch=master)

This is a operator which handles rabbitmq resources via kubernetes custom resource definitions.

# Supported CRDs

- Users
- Permissions
- Exchanges
- Queues
- Bindings
- Shovles

# Getting Started

## Install the CRDs

```console
# *** This is for GKE Regular Channel - k8s 1.16 -> Adjust based on your cloud or storage options
kubectl apply -f https://raw.githubusercontent.com/datastax/cass-operator/v1.4.1/docs/user/cass-operator-manifests-v1.16.yaml
kubectl create -f https://raw.githubusercontent.com/datastax/cass-operator/v1.4.1/operator/k8s-flavors/gke/storage.yaml
kubectl -n cass-operator create -f https://raw.githubusercontent.com/datastax/cass-operator/v1.4.1/operator/example-cassdc-yaml/cassandra-3.11.x/example-cassdc-minimal.yaml
```

## Install the operator

```console
kubectl apply -f https://raw.githubusercontent.com/datastax/cass-operator/v1.4.1/docs/user/cass-operator-manifests-$K8S_VER.yaml
```


# Future Plan

- Support all rabbitmq rest resources
- support rabbitmq cluster setup 
