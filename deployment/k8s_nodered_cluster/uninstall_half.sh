#!/bin/bash

# Delete half deployments and services in the namespace
kubectl delete -f openiiot-nodered-deployment-11.yaml
kubectl delete -f openiiot-nodered-service-11.yaml
kubectl delete -f openiiot-nodered-deployment-12.yaml
kubectl delete -f openiiot-nodered-service-12.yaml
kubectl delete -f openiiot-nodered-deployment-13.yaml
kubectl delete -f openiiot-nodered-service-13.yaml
kubectl delete -f openiiot-nodered-deployment-14.yaml
kubectl delete -f openiiot-nodered-service-14.yaml
kubectl delete -f openiiot-nodered-deployment-15.yaml
kubectl delete -f openiiot-nodered-service-15.yaml
kubectl delete -f openiiot-nodered-deployment-16.yaml
kubectl delete -f openiiot-nodered-service-16.yaml
kubectl delete -f openiiot-nodered-deployment-17.yaml
kubectl delete -f openiiot-nodered-service-17.yaml
kubectl delete -f openiiot-nodered-deployment-18.yaml
kubectl delete -f openiiot-nodered-service-18.yaml
kubectl delete -f openiiot-nodered-deployment-19.yaml
kubectl delete -f openiiot-nodered-service-19.yaml
kubectl delete -f openiiot-nodered-deployment-20.yaml
kubectl delete -f openiiot-nodered-service-20.yaml

# Optionally, wait for a few seconds to ensure that all resources are properly terminated
echo "Waiting for resources to terminate..."
sleep 5

echo "Uninstallation completed successfully."
