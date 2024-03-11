kubectl delete -f deployment/k8s/openiiot-server-deployment.yaml
docker rmi openiiot_server:1.0.1
echo "stop old deployment successfully"

docker load -i deployment/binary/openiiot_server.tar
kubectl apply -f deployment/k8s/openiiot-server-deployment.yaml
echo "start new server deployment successfully"
