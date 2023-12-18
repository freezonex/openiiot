## How to test serser api inside k8s:
```shell
kubectl exec -it test -n openiiot -- curl http://openiiot-server-service.openiiot:8085/ping
```
## Check if there are nfs volume
```shell
showmount -e localhost
```
if there are mount point, the result will be
```shell
Export list for localhost:
/volumes/nfs *
```
## Image update
if you have a new image file and use "docker load -i new_iamge", the old image with same name and tag will be override  
after that, since no deployment.yaml update, you can use
```shell
kubectl rollout restart deployment <deployment-name>
```
to make rollout restart to use new image  
for example
```shell
kubectl rollout restart deployment/openiiot-nodered -n openiiot
```
