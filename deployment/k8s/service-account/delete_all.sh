# check serviceaccount status
kubectl describe serviceaccount admin-user -n kube-system
kubectl describe secret admin-user-manual-token -n kube-system
kubectl describe secret admin-user-manual-token -n kube-system

# delete serviceaccount
kubectl delete clusterrolebinding admin-user -n kube-system
kubectl delete secret admin-user-manual-token -n kube-system
kubectl delete serviceaccount admin-user -n kube-system