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

echo "Installation environment check passed"