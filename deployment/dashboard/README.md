# kubernets dashboard

## Install Kubernetes-Dashboard

```bash
# Add dashboard repo to your helm
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
# Create dashboard namespace
kubectl create ns dashboard
# Install dashboard to the last namespace
helm install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard -n dashboard
# Apply role and account config file
kubectl apply -f cluster_role_binding.yaml
kubectl apply -f service_account.yaml
# Get the login token
kubectl -n dashboard create token admin-user > token.txt
```

## Visit Kubernetes-Dashboard

```bash
export POD_NAME=$(kubectl get pods  \
        -n dashboard \
        -l "app.kubernetes.io/name=kubernetes-dashboard,app.kubernetes.io/instance=kubernetes-dashboard" \
        -o jsonpath="{.items[0].metadata.name}")
echo Visit the dashboard https://127.0.0.1:8443/ by your token:
cat token.txt
kubectl -n dashboard port-forward $POD_NAME 8443:8443
```
