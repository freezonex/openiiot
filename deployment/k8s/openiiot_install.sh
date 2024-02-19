# docker load -i all images
kubectl create -f create-service-account.yaml
kubectl create -f role-binding.yaml
kubectl describe secret $secret_name -n kube-system | grep token