#!/bin/bash

# Set Docker Registry Prefix
DOCKER_REGISTRY_PREFIX="registry:5000/"
TAG="1.0.0"

# Iterate over each tar file in the current directory
for file in *.tar; do
    if [[ -f "$file" ]]; then
        # Extract the image name from the filename (assuming format is openiiot_servicename.tar)
        service_name=$(basename "$file" .tar | sed 's/^openiiot_//')

        # Load the image from the tar file
        echo "Loading image from $file..."
        docker load -i "$file"

        # Tag the image with the new prefix
        original_image="openiiot_${service_name}:${TAG}"
        new_image="${DOCKER_REGISTRY_PREFIX}openiiot_${service_name}:${TAG}"
        echo "Tagging image $original_image as $new_image..."
        docker tag "$original_image" "$new_image"

        # Push the image to the registry
        echo "Pushing image $new_image..."
        docker push "$new_image"
    fi
done

echo "All images processed."
