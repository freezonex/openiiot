#!/bin/bash

# Function to wait for a job to complete
wait_for_job() {
    local job_name=$1
    local namespace=$2
    echo "Waiting for job $job_name to complete..."
    while true; do
        if kubectl get job $job_name -n $namespace | grep -q '1/1'; then
            echo "Job $job_name completed successfully."
            break
        else
            echo "Waiting for job $job_name to complete..."
            sleep 10
        fi
    done
}

# Check if the namespace "openiiot" does not exist
if ! kubectl get namespace openiiot > /dev/null 2>&1; then
    echo "Namespace 'openiiot' does not exist. Please use "kubectl create namespace openiiot" to create it before running this script."
    exit 1
fi

# Apply pv and pvc YAML files
kubectl apply -f manifest/openiiot-nodered-pv.yaml
kubectl apply -f manifest/openiiot-nodered-pvc.yaml
kubectl apply -f manifest/openiiot-grafana-pv.yaml
kubectl apply -f manifest/openiiot-grafana-pvc.yaml
kubectl apply -f manifest/openiiot-emqx-pv.yaml
kubectl apply -f manifest/openiiot-emqx-pvc.yaml
kubectl apply -f manifest/openiiot-tdengine-data-pv.yaml
kubectl apply -f manifest/openiiot-tdengine-data-pvc.yaml
kubectl apply -f manifest/openiiot-tdengine-log-pv.yaml
kubectl apply -f manifest/openiiot-tdengine-log-pvc.yaml
kubectl apply -f manifest/openiiot-mysql-pv.yaml
kubectl apply -f manifest/openiiot-mysql-pvc.yaml

# Apply the jobs
kubectl apply -f manifest/install-nodered-plugins-job.yaml
kubectl apply -f manifest/install-grafana-plugins-job.yaml

# Wait for both jobs to complete
wait_for_job "install-nodered-plugins" "openiiot"
wait_for_job "install-grafana-plugins" "openiiot"

# Apply deployment and service YAML files
kubectl apply -f manifest/openiiot-server-deployment.yaml
kubectl apply -f manifest/openiiot-server-service.yaml
kubectl apply -f manifest/openiiot-consolemanager-deployment.yaml
kubectl apply -f manifest/openiiot-consolemanager-service.yaml
kubectl apply -f manifest/openiiot-nodered-deployment.yaml
kubectl apply -f manifest/openiiot-nodered-service.yaml
kubectl apply -f manifest/openiiot-grafana-deployment.yaml
kubectl apply -f manifest/openiiot-grafana-service.yaml
kubectl apply -f manifest/openiiot-emqx-deployment.yaml
kubectl apply -f manifest/openiiot-emqx-service.yaml
kubectl apply -f manifest/openiiot-tdengine-deployment.yaml
kubectl apply -f manifest/openiiot-tdengine-service.yaml
kubectl apply -f manifest/openiiot-mysql-deployment.yaml
kubectl apply -f manifest/openiiot-mysql-service.yaml
kubectl apply -f manifest/openiiot-pma-deployment.yaml
kubectl apply -f manifest/openiiot-pma-service.yaml

echo "Installation completed successfully."
