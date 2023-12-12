#!/bin/bash

# Delete all deployments and services in the namespace
kubectl delete deployments,services --all -n openiiot

# Optionally, wait for a few seconds to ensure that all resources are properly terminated
echo "Waiting for resources to terminate..."
sleep 10

# Delete the namespace
kubectl delete namespace openiiot

echo "Uninstallation completed successfully."
