#!/bin/bash

# Define port range
START_PORT=30050
END_PORT=30070

check_ports() {
    for (( port = START_PORT; port <= END_PORT; port++ )); do
        if lsof -i:"$port" > /dev/null 2>&1; then
            echo "Port $port is occupied. Exiting installation."
            exit 1
        fi
    done
}

# Check if any port in the range is occupied
check_ports


# Check if the namespace "openiiot" already exists
if kubectl get namespace openiiot-nodered > /dev/null 2>&1; then
    echo "Namespace 'openiiot-nodered' already exists."
    exit 1
fi

# Create the namespace "openiiot"
kubectl create namespace openiiot-nodered

# Apply deployment and service YAML files
kubectl apply -f openiiot-nodered-deployment-1.yaml
kubectl apply -f openiiot-nodered-service-1.yaml
kubectl apply -f openiiot-nodered-deployment-2.yaml
kubectl apply -f openiiot-nodered-service-2.yaml
kubectl apply -f openiiot-nodered-deployment-3.yaml
kubectl apply -f openiiot-nodered-service-3.yaml
kubectl apply -f openiiot-nodered-deployment-4.yaml
kubectl apply -f openiiot-nodered-service-4.yaml
kubectl apply -f openiiot-nodered-deployment-5.yaml
kubectl apply -f openiiot-nodered-service-5.yaml
kubectl apply -f openiiot-nodered-deployment-6.yaml
kubectl apply -f openiiot-nodered-service-6.yaml
kubectl apply -f openiiot-nodered-deployment-7.yaml
kubectl apply -f openiiot-nodered-service-7.yaml
kubectl apply -f openiiot-nodered-deployment-8.yaml
kubectl apply -f openiiot-nodered-service-8.yaml
kubectl apply -f openiiot-nodered-deployment-9.yaml
kubectl apply -f openiiot-nodered-service-9.yaml
kubectl apply -f openiiot-nodered-deployment-10.yaml
kubectl apply -f openiiot-nodered-service-10.yaml
kubectl apply -f openiiot-nodered-deployment-11.yaml
kubectl apply -f openiiot-nodered-service-11.yaml
kubectl apply -f openiiot-nodered-deployment-12.yaml
kubectl apply -f openiiot-nodered-service-12.yaml
kubectl apply -f openiiot-nodered-deployment-13.yaml
kubectl apply -f openiiot-nodered-service-13.yaml
kubectl apply -f openiiot-nodered-deployment-14.yaml
kubectl apply -f openiiot-nodered-service-14.yaml
kubectl apply -f openiiot-nodered-deployment-15.yaml
kubectl apply -f openiiot-nodered-service-15.yaml
kubectl apply -f openiiot-nodered-deployment-16.yaml
kubectl apply -f openiiot-nodered-service-16.yaml
kubectl apply -f openiiot-nodered-deployment-17.yaml
kubectl apply -f openiiot-nodered-service-17.yaml
kubectl apply -f openiiot-nodered-deployment-18.yaml
kubectl apply -f openiiot-nodered-service-18.yaml
kubectl apply -f openiiot-nodered-deployment-19.yaml
kubectl apply -f openiiot-nodered-service-19.yaml
kubectl apply -f openiiot-nodered-deployment-20.yaml
kubectl apply -f openiiot-nodered-service-20.yaml

echo "Installation completed successfully."
