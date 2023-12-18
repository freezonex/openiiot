#!/bin/bash

# Define the base directory
BASE_DIR="/volumes/nfs/openiiot"

# Array of subdirectories
SUBDIRS=("nodered_data" "grafana_data" "tdengine_data" "tdengine_log" "emqx_data" "mysql_data")

# Check if the base directory exists, if not create it
if [ ! -d "$BASE_DIR" ]; then
    echo "Creating base directory: $BASE_DIR"
    mkdir -p "$BASE_DIR"
fi

# Loop through and create each subdirectory
for subdir in "${SUBDIRS[@]}"; do
    dir_path="$BASE_DIR/$subdir"
    if [ ! -d "$dir_path" ]; then
        echo "Creating directory: $dir_path"
        mkdir -p "$dir_path"
        chmod 777 "$dir_path"
    else
        echo "Directory already exists: $dir_path"
    fi
done

echo "Directory structure creation complete."
