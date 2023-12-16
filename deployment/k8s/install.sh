#!/bin/bash

# Array of ports to check
PORTS=(30085 30080 30081 31880 31893 32000 30883 30083 30084 31883 31083 32030 32041 32306)

check_ports() {
    for port in "${PORTS[@]}"; do
        if lsof -i:"$port" > /dev/null 2>&1; then
            echo "Port $port is occupied. Exiting installation."
            exit 1
        fi
    done
}

# Check if any port is occupied
check_ports

# Check if the namespace "openiiot" already exists
if kubectl get namespace openiiot > /dev/null 2>&1; then
    echo "Namespace 'openiiot' already exists. Exiting installation."
    exit 1
fi

# Create the namespace "openiiot"
kubectl create namespace openiiot

# Apply deployment and service YAML files
kubectl apply -f openiiot-server-deployment.yaml
kubectl apply -f openiiot-server-service.yaml
kubectl apply -f openiiot-web-deployment.yaml
kubectl apply -f openiiot-web-service.yaml
kubectl apply -f openiiot-nodered-deployment.yaml
kubectl apply -f openiiot-nodered-service.yaml
kubectl apply -f openiiot-grafana-deployment.yaml
kubectl apply -f openiiot-grafana-service.yaml
kubectl apply -f openiiot-emqx-deployment.yaml
kubectl apply -f openiiot-emqx-service.yaml
kubectl apply -f openiiot-tdengine-deployment.yaml
kubectl apply -f openiiot-tdengine-service.yaml
kubectl apply -f openiiot-mysql-deployment.yaml
kubectl apply -f openiiot-mysql-service.yaml

echo "Installation completed successfully."
