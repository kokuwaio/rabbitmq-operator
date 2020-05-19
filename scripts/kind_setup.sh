#!/usr/bin/env bash

set -e

kind create cluster --image "kindest/node:v1.15.0" --config ../.github/kind-config.yaml --name development

echo
echo "--------------------"
echo "Setting kube-context"
echo "--------------------"
export KUBECONFIG="$(kind get kubeconfig-path --name="development")"
kubectl cluster-info