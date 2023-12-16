## K8S version confirm
Since SupOS use 1.18 k8s cluter, so need modify a lot of files
## Create Helm Chart
```shell
helm create nodered-cluster-chart
```

## Customize the Chart
```shell
cd nodered-cluster-chart
```
Inside the chart directory, you will find several files and directories. The main ones to focus on are:

* Chart.yaml: The chart's metadata file.
* values.yaml: The default values for your chart's parameters.
* templates/: A directory where your Kubernetes YAML files (templates) will reside.

## Define Templates and Values
1. Modify `values.yaml`: Define the values that will be different for each deployment and service. 
2. Create `templates/deployment.yaml` and `templates/service.yaml`: Use Helm templating syntax to parameterize these files.

## Create the namespace "openiiot-nodered-cluster"
kubectl create namespace openiiot-nodered-cluster

## Install the Chart
Once your chart is ready, you can install it using the helm install command.  
This will create the defined number of pods and services based on your templates and values.
```shell
helm install nodered-cluster-chart ./nodered-cluster-chart
```

## Uninstall the Chart
```shell
helm uninstall nodered-cluster-chart
```

## Upgrade the Chart
```shell
helm upgrade nodered-cluster-chart ./nodered-cluster-chart
```