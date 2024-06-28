#!/bin/bash

# List of components
components=("consolemanager" "emqx" "grafana" "mysql" "nodered" "server" "tdengine")

# Loop through each component and execute the docker commands
for component in "${components[@]}"; do
    echo "Loading $component..."
    
    # Load the Docker image
    docker load -i binary/openiiot_${component}.tar || true
    
    # Tag the Docker image
    docker tag openiiot_${component}:1.0.0 localhost:5000/openiiot_${component}:1.0.0 || true
    
    # Push the Docker image to the local registry
    docker push localhost:5000/openiiot_${component}:1.0.0 || true
    
    echo "$component loaded."
done

echo "All images loaded."
