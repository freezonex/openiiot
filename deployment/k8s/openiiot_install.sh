#!/bin/bash

# Check if the namespace "openiiot" does not exist
if ! kubectl get namespace openiiot > /dev/null 2>&1; then
    echo "Namespace 'openiiot' does not exist. Please use "kubectl create namespace openiiot" to create it before running this script."
    exit 1
fi

# docker load -i all images
kubectl create -f create-service-account.yaml
kubectl create -f role-binding.yaml
kubectl describe secret $secret_name -n kube-system | grep token