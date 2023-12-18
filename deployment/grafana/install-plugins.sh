#!/bin/bash

# Define the list of plugins
PLUGINS=("grafana-mqtt-datasource" "satellogic-3d-globe-panel" "cloudspout-button-panel" "grafana-iot-twinmaker-app" "tdengine-datasource")

# Path to the Grafana plugins directory
PLUGIN_DIR="/var/lib/grafana/plugins"

# Install a plugin if it's not already installed
for plugin in "${PLUGINS[@]}"; do
    if [ ! -d "$PLUGIN_DIR/$plugin" ]; then
        echo "Installing plugin: $plugin"
        grafana cli plugins install $plugin
    else
        echo "Plugin already installed: $plugin"
    fi
done

echo "Plugin installation process completed."
