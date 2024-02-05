# Run this script when install openiiot standalone

# Update system and install dependencies
sudo apt-get update && sudo apt-get install -y apt-transport-https ca-certificates curl

# Install Docker
sudo apt-get install -y docker.io

# Install kubeadm, kubelet and kubectl
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | sudo apt-key add -
echo "deb https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl

# Prevent Automatic Updates
sudo apt-mark hold kubelet kubeadm kubectl

# Initialize the Kubernetes cluster
sudo kubeadm init --pod-network-cidr=192.168.0.0/16

# Configure kubectl
mkdir -p $HOME/.kube
sudo cp -i /etc/kubernetes/admin.conf $HOME/.kube/config
sudo chown $(id -u):$(id -g) $HOME/.kube/config

# Install network plug-in
kubectl apply -f https://docs.projectcalico.org/manifests/calico.yaml

# Convert node to single node cluster, first line is old k8s with master taint, newer k8s use contol plane
# kubectl taint nodes --all node-role.kubernetes.io/master-
kubectl taint nodes --all node-role.kubernetes.io/control-plane:NoSchedule-


# Verify installation
kubectl get nodes
kubectl get pods --all-namespaces
