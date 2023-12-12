#!/bin/bash

# Array of ports to check
PORTS=(30085 30080 31880 31893 32000 30883 30083 30084 31883 31083 32030 32041 32306)

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
    echo "Namespace 'openiiot' already exists."
    exit 1
fi

echo "Installation environment check passed"