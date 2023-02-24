PROJECT_ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	helm repo add bitnami https://charts.bitnami.com/bitnami

install-cluster:
	-kind delete cluster --name douyin
	kind create cluster --config deployment/cluster/douyin-cluster.yaml

### 中间件 开始

install-minio:
	-kubectl delete ns minio
	helm repo add minio https://charts.min.io/
	kubectl create ns minio
	helm install minio minio/minio -f deployment/minio/minio.yaml -n minio

install-etcd:
	-kubectl delete ns etcd
	kubectl create ns etcd
	helm install etcd bitnami/etcd --set replicaCount=2 -n etcd --set auth.rbac.create=false

install-kafka:
	-kubectl delete ns kafka
	kubectl create ns kafka
	helm install kafka bitnami/kafka -n kafka

install-redis:
	-kubectl delete ns redis
	kubectl create ns redis
	helm install redis bitnami/redis -n redis --set auth.password='redispwd123'

### 中间件 结束
### 微服务 开始

install-minio-client:
	docker build -f ${PROJECT_ROOT}/cmd/minio-client/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/minio-client:nightly
	-kubectl delete ns minio-client
	-kind load docker-image douyin/minio-client:nightly --name douyin
	kubectl create ns minio-client
	kubectl apply -f deployment/minio-client/minio-client.yaml


install-user:
	docker build -f ${PROJECT_ROOT}/cmd/user/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/user:nightly
	-kubectl delete ns user
	-kind load docker-image douyin/user:nightly --name douyin
	kubectl create ns user
	kubectl apply -f deployment/user/user.yaml

install-comment:
	docker build -f ${PROJECT_ROOT}/cmd/comment/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/comment:nightly
	-kubectl delete ns comment
	-kind load docker-image douyin/comment:nightly --name douyin
	kubectl create ns comment
	kubectl apply -f deployment/comment/comment.yaml

install-mq:
	docker build -f ${PROJECT_ROOT}/cmd/mq/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/mq:nightly
	-kubectl delete ns mq
	-kind load docker-image douyin/mq:nightly --name douyin
	kubectl create ns mq
	kubectl apply -f deployment/mq/mq.yaml

install-favorite:
	docker build -f ${PROJECT_ROOT}/cmd/favorite/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/favorite:nightly
	-kubectl delete ns favorite
	-kind load docker-image douyin/favorite:nightly --name douyin
	kubectl create ns favorite
	kubectl apply -f deployment/favorite/favorite.yaml

install-video:
	docker build -f ${PROJECT_ROOT}/cmd/video/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/video:nightly
	-kubectl delete ns video
	-kind load docker-image douyin/video:nightly --name douyin
	kubectl create ns video
	kubectl apply -f deployment/video/video.yaml

install-message:
	docker build -f ${PROJECT_ROOT}/cmd/message/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/message:nightly
	-kubectl delete ns message
	-kind load docker-image douyin/message:nightly --name douyin
	kubectl create ns message
	kubectl apply -f deployment/message/message.yaml

install-follow:
	docker build -f ${PROJECT_ROOT}/cmd/follow/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/follow:nightly
	-kubectl delete ns follow
	-kind load docker-image douyin/follow:nightly --name douyin
	kubectl create ns follow
	kubectl apply -f deployment/follow/follow.yaml

install-gateway:
	docker build -f ${PROJECT_ROOT}/cmd/gateway/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/gateway:nightly
	-kubectl delete ns gateway
	-kind load docker-image douyin/gateway:nightly --name douyin
	kubectl create ns gateway
	kubectl apply -f deployment/gateway/gateway.yaml

### 微服务 结束

# Deploy nfs -> Declare nfs pv -> Inject sql scheme -> Deploy mysql
nfs-init-service:
	kubectl apply -f deployment/nfs/nfs-deploy.yaml
	kubectl apply -f deployment/nfs/nfs-pvx.yaml
mysql-init-service: nfs-init-service
	kubectl apply -f deployment/mysql/mysql-scheme.yaml
	kubectl apply -f deployment/mysql/mysql-deploy.yaml
mysql-regenerate-codes:
	go run cmd/mysql-gen/gen.go

forward-minio-console:
	kubectl port-forward -n minio svc/minio-console 9001:9001

forward-gateway:
	kubectl port-forward -n gateway svc/gateway 8888: --address='0.0.0.0'

forward-minio:
	kubectl port-forward -n minio svc/minio 9000:9000 --address='0.0.0.0'
