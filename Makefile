# 部署所有服务 todo:build image
install:
	@echo 开始清空Kubernetes-douyin
	kind delete cluster --name douyin
	@echo 开始部署Kubernetes
	kind create cluster --config deployment/cluster/douyin-cluster.yaml
	@echo 开始部署MinIO
	helm repo add minio https://charts.min.io/
	helm install minio minio/minio -f deployment/minio/minio.yaml -n minio --create-namespace

# 格式化所有代码
fmt:
	@echo 开始格式化所有go代码
	go fmt ./*

# 转发minio的api服务
forward-minio-api:
	@echo 正在转发MinIO的API服务,请通过127.0.0.1:9000访问
	kubectl port-forward -n minio svc/minio 9000:9000

# 转发minio的console服务
forward-minio-console:
	@echo 正在转发MinIO的console服务,请通过127.0.0.1:9001访问
	kubectl port-forward -n minio svc/minio-console 9001:9001

# todo:forward dashboard