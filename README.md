# Rabbitmq Operator

[![Gitpod ready-to-code](https://img.shields.io/badge/Gitpod-ready--to--code-blue?logo=gitpod)](https://gitpod.io/from-referrer/)
![CI build and Deploy](https://github.com/kokuwaio/rabbitmq-operator/workflows/CI%20build%20and%20Deploy/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/kokuwaio/rabbitmq-operator/badge.svg?branch=master)](https://coveralls.io/github/kokuwaio/rabbitmq-operator?branch=master)

This is a operator which handles rabbitmq resources via kubernetes custom resource definitions. To manage rabbitmq clsuters you can use for example the [rabbitmq operator](https://github.com/rabbitmq/cluster-operator).

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
kubectl apply -f https://raw.githubusercontent.com/kokuwaio/rabbitmq-operator/master/config/crd/bases/crds.yaml
```

## Install the operator

```console
kubectl apply -f https://raw.githubusercontent.com/kokuwaio/rabbitmq-operator/master/config/manager/manager.yaml
```

## Using the operator

### creating a cluster reference

```yaml
apiVersion: rabbitmq.kokuwa.io/v1beta1
kind: RabbitmqCluster
metadata:
  name: rabbitmqcluster-sample
spec:
  # URL for the admin api
  host: http://rabbitmq.default.svc:15672
  # username for the rabbitmq cluster
  # will only be used if secretRef.userKey is empty
  user: admin
  secretRef:
    # name of the secret
    name: rabbitmq
    # namespace where the secret is located
    # the crd is not namespace scoped, so the namespace is a required field
    namespace: default
    # key for the password if empty password will be used
    passwordKey: rabbitmq-password
    # key for the suer can be empty
    userKey: rabbitmq-user
```

### creating a user

password secret:

```yaml

```

user crd:

```yaml
apiVersion: rabbitmq.kokuwa.io/v1beta1
kind: RabbitmqUser
metadata:
  name: rabbitmquser-sample
spec:
  # name for the user
  name: my-examle-user
  # tags for the user
  tags: administrator
  secretRef:
    name: my-secret
    namespace: default
    key: password
  clusterRef:
    name: rabbitmqcluster-sample
```

### creating a queue


# Future Plan

- Support all rabbitmq rest resources
