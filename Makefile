PROJECT_ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

# 部署所有服务 todo:build image
install:
	@echo 开始清空Kubernetes-douyin
	kind delete cluster --name douyin
	@echo 开始部署Kubernetes
	kind create cluster --config deployment/cluster/douyin-cluster.yaml
	@echo 开始部署MinIO
	helm repo add minio https://charts.min.io/
	kubectl create ns minio
	helm install minio minio/minio -f deployment/minio/minio.yaml -n minio

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



# 从编译到安装一个Demo的全过程
# 1.编译要用的image
# 2.用kind load,把编译好的image载入到你的节点里
# 3.用kubectl通过yaml创建Deployment和Service
# 4.转发出demo提供服务的端口(80),然后访问localhost:8000/hello
buildx-demo:
	docker buildx build --platform=linux/arm64 -f ${PROJECT_ROOT}/cmd/demo/Dockerfile \
	--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
	-t douyin/demo:nightly

install-demo:
	-kubectl delete ns demo
	kind load docker-image douyin/demo:nightly --name douyin
	-kubectl create ns demo
	kubectl apply -f deployment/demo/demo.yaml

forward-demo:
	kubectl port-forward -n demo service/demo 8000:80