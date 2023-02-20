PROJECT_ROOT := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

# Todo:build all image
init:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go install google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
	go install google.golang.org/grpc/cmd/protoc-gen-go-grpc
	helm repo add bitnami https://charts.bitnami.com/bitnami

install-cluster:
	kind delete cluster --name douyin
	kind create cluster --config deployment/cluster/douyin-cluster.yaml

# Deploy MinIO Service
install-minio:
	-kubectl delete ns minio
	helm repo add minio https://charts.min.io/
	kubectl create ns minio
	helm install minio minio/minio -f deployment/minio/minio.yaml -n minio

# Forward MinIO's console service(web)
forward-minio-console:
	kubectl port-forward -n minio svc/minio-console 9001:9001


# Deploy Etcd service, you must install it before use rpc
install-etcd:
	-kubectl delete ns etcd
	helm repo add etcd https://charts.bitnami.com/bitnami
	kubectl create ns etcd
	helm install etcd etcd/etcd --set replicaCount=2 -n etcd --set auth.rbac.create=false



install-kafka:
	-kubectl delete ns kafka
	helm repo add bitnami https://charts.bitnami.com/bitnami
	kubectl create ns kafka
	helm install kafka bitnami/kafka -n kafka --set replicaCount=2

forward-kafka:
	kubectl port-forward -n kafka svc/kafka 9092:9092

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


build-user:
	docker build -f ${PROJECT_ROOT}/cmd/user/rpc/Dockerfile \
		--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
		-t douyin/user-rpc:nightly
	docker build -f ${PROJECT_ROOT}/cmd/user/api/Dockerfile \
    	--build-arg PROJECT_ROOT="${PROJECT_ROOT}" ${PROJECT_ROOT} \
    	-t douyin/user-api:nightly

install-user: build-user
	-kubectl delete ns user
	kind load docker-image douyin/user-rpc:nightly --name douyin
	kind load docker-image douyin/user-api:nightly --name douyin
	kubectl create ns user
	kubectl apply -f deployment/user/user.yaml

