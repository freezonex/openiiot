#!/bin/bash

# Delete all deployments and services in the namespace
kubectl delete deployments,services --all -n openiiot-nodered

# Optionally, wait for a few seconds to ensure that all resources are properly terminated
echo "Waiting for resources to terminate..."
sleep 5

# Delete the namespace
kubectl delete namespace openiiot-nodered

echo "Uninstallation completed successfully."
