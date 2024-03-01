#!/bin/bash

# Check if the namespace "openiiot" does not exist
if ! kubectl get namespace openiiot > /dev/null 2>&1; then
    echo "Namespace 'openiiot' does not exist. Please use "kubectl create namespace openiiot" to create it before running this script."
    exit 1
fi

# docker load -i all images
kubectl create -f create-service-account.yaml
kubectl create -f role-binding.yaml

# Set the ServiceAccount name and namespace
SERVICE_ACCOUNT_NAME="admin-user"
NAMESPACE="kube-system"

# when k8s >= v1.24.0, secret will not be automatically generated once ServiceAccount created, needs create manually
# Get the Secret name associated with the ServiceAccount
SECRET_NAME=$(kubectl get serviceaccount $SERVICE_ACCOUNT_NAME -n $NAMESPACE -o jsonpath="{.secrets[0].name}")

# Get and print Token
echo "Token for ServiceAccount $SERVICE_ACCOUNT_NAME in $NAMESPACE namespace:"
kubectl get secret $SECRET_NAME -n $NAMESPACE -o jsonpath="{.data.token}" | base64 --decode
echo

kubectl describe secret $secret_name -n kube-system | grep token