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
kubectl apply -f openiiot-nodered-pv.yaml
kubectl apply -f openiiot-nodered-pvc.yaml
kubectl apply -f openiiot-grafana-pv.yaml
kubectl apply -f openiiot-grafana-pvc.yaml
kubectl apply -f openiiot-emqx-pv.yaml
kubectl apply -f openiiot-emqx-pvc.yaml
kubectl apply -f openiiot-tdengine-data-pv.yaml
kubectl apply -f openiiot-tdengine-data-pvc.yaml
kubectl apply -f openiiot-tdengine-log-pv.yaml
kubectl apply -f openiiot-tdengine-log-pvc.yaml
kubectl apply -f openiiot-mysql-pv.yaml
kubectl apply -f openiiot-mysql-pvc.yaml

# Apply the jobs
kubectl apply -f install-nodered-plugins-job.yaml
kubectl apply -f install-grafana-plugins-job.yaml

# Wait for both jobs to complete
wait_for_job "install-nodered-plugins" "openiiot"
wait_for_job "install-grafana-plugins" "openiiot"

# Apply deployment and service YAML files
kubectl apply -f openiiot-server-deployment.yaml
kubectl apply -f openiiot-server-service.yaml
kubectl apply -f openiiot-consolemanager-deployment.yaml
kubectl apply -f openiiot-consolemanager-service.yaml
kubectl apply -f openiiot-web-deployment.yaml
kubectl apply -f openiiot-web-service.yaml
kubectl apply -f openiiot-nodered-deployment.yaml
kubectl apply -f openiiot-nodered-service.yaml
kubectl apply -f openiiot-grafana-deployment.yaml
kubectl apply -f openiiot-grafana-service.yaml
kubectl apply -f openiiot-emqx-deployment.yaml
kubectl apply -f openiiot-emqx-service.yaml
kubectl apply -f openiiot-tdengine-deployment.yaml
kubectl apply -f openiiot-tdengine-service.yaml
kubectl apply -f openiiot-mysql-deployment.yaml
kubectl apply -f openiiot-mysql-service.yaml

echo "Installation completed successfully."
