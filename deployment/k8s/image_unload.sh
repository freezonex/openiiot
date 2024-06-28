#!/bin/bash

# List of components
components=("consolemanager" "emqx" "grafana" "mysql" "nodered" "server" "tdengine")

# Loop through each component and execute the docker commands to remove the images
for component in "${components[@]}"; do
    echo "Removing $component..."
    
    # Remove the tagged Docker image
    docker rmi localhost:5000/openiiot_${component}:1.0.0 || true
    
    # Remove the original Docker image
    docker rmi openiiot_${component}:1.0.0 || true
    
    echo "$component removed."
done

echo "All images removed."
