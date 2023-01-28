## 安装Dashboard
```bash
# 添加Dashboard到库
helm repo add kubernetes-dashboard https://kubernetes.github.io/dashboard/
# 创建dashboard namespace(windows的helm好像没有--create-namespace)
kubectl create ns dashboard
# 安装dashboard到dashboard namespace
helm install kubernetes-dashboard kubernetes-dashboard/kubernetes-dashboard -n dashboard
# 载入dashboard目录下的两个配置文件(创建权限和account)
kubectl apply -f cluster_role_binding.yaml
kubectl apply -f service_account.yaml
# 获得token(后续直接从dashboard/token.txt文件中拿token)
kubectl -n dashboard create token admin-user > token.txt
```

## 访问Dashboard
```bash
# 拿到你创建的Dashboard的pod的名称
kubectl get pod -n dashboard | grep "kubernetes-dashboard"
# 把pod内部的端口转发出来(把pod名称改成你自己的)
kubectl port-forward -n dashboard your_pod_name 8443:8443
# 访问http://127.0.0.1:8443,使用你上一步创建的token来登录
```