#!/bin/bash

# Change to the Node-RED data directory
cd /data

# Install plugins defined in package.json
echo "Installing Node-RED plugins..."
npm install

echo "Node-RED plugin installation complete."
