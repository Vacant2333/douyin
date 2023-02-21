PROJECT_ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

# Todo:build all image
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	helm repo add bitnami https://charts.bitnami.com/bitnami

install-cluster:
	-kind delete cluster --name douyin
	kind create cluster --config deployment/cluster/douyin-cluster.yaml

# Deploy MinIO Service
install-minio:
	-kubectl delete ns minio
	helm repo add minio https://charts.min.io/
	kubectl create ns minio
	helm install minio minio/minio -f deployment/minio/minio.yaml -n minio

# Deploy Etcd service, you must install it before use rpc
install-etcd:
	-kubectl delete ns etcd
	kubectl create ns etcd
	helm install etcd bitnami/etcd --set replicaCount=2 -n etcd --set auth.rbac.create=false

install-kafka:
	-kubectl delete ns kafka
	kubectl create ns kafka
	helm install kafka bitnami/kafka -n kafka --set replicaCount=2

install-redis:
	-kubectl delete ns redis
	kubectl create ns redis
	helm install redis bitnami/redis -n redis --set auth.password='redispwd123'




install-user:
	docker build -f ${PROJECT_ROOT}/cmd/user/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/user:nightly
	-kubectl delete ns user
	kind load docker-image douyin/user:nightly --name douyin
	kubectl create ns user
	kubectl apply -f deployment/user/user.yaml

install-comment:
	docker build -f ${PROJECT_ROOT}/cmd/comment/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/comment:nightly
	-kubectl delete ns comment
	kind load docker-image douyin/comment:nightly --name douyin
	kubectl create ns comment
	kubectl apply -f deployment/comment/comment.yaml

install-favorite:
	docker build -f ${PROJECT_ROOT}/cmd/favorite/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/favorite:nightly
	-kubectl delete ns favorite
	kind load docker-image douyin/favorite:nightly --name douyin
	kubectl create ns favorite
	kubectl apply -f deployment/favorite/favorite.yaml

install-video:
	docker build -f ${PROJECT_ROOT}/cmd/video/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/video:nightly
	-kubectl delete ns video
	kind load docker-image douyin/video:nightly --name douyin
	kubectl create ns video
	kubectl apply -f deployment/video/video.yaml

install-gateway:
	docker build -f ${PROJECT_ROOT}/cmd/gateway/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/gateway:nightly
	-kubectl delete ns gateway
	kind load docker-image douyin/gateway:nightly --name douyin
	kubectl create ns gateway
	kubectl apply -f deployment/gateway/gateway.yaml





# todo: delete deployment/comment/* != comment.yaml






# Build proto
build-proto-minio-client:
	cd pkg/minio-client && goctl rpc protoc ./proto/minio-client.proto --go_out=./types --go-grpc_out=./types --zrpc_out=.

# Deploy nfs -> Declare nfs pv -> Inject sql scheme -> Deploy mysql
nfs-init-service:
	kubectl apply -f deployment/nfs/nfs-deploy.yaml
	kubectl apply -f deployment/nfs/nfs-pvx.yaml
mysql-init-service: nfs-init-service
	kubectl apply -f deployment/mysql/mysql-scheme.yaml
	kubectl apply -f deployment/mysql/mysql-deploy.yaml
mysql-regenerate-codes:
	go run cmd/mysql-gen/gen.go

forward-gateway:
	kubectl port-forward -n gateway svc/gateway 8888: --address='0.0.0.0'