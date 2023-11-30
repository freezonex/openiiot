#!/bin/bash

# Define the image prefix to search for
IMAGE_PREFIX="openiiot"

# Find all images that match the prefix and delete them
docker images | grep "$IMAGE_PREFIX" | awk '{print $3}' | uniq | xargs -r docker rmi -f
