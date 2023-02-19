goctl model mysql datasource -url="user:password@tcp(127.0.0.1:3306)/database" -table="*"  -dir="./model" goctl model mysql datasource -url="user:password@tcp(127.0.0.1:3306)/database" -table="*"  -dir="./model"

kubectl apply -f auth.yaml

kubectl get sa -n douyin-comment

goctl docker -go .\usercomment.go .

docker build -t douyin-comment-api:v1 .

kind create cluster --image kindest/node:v1.25.3 --config deployment/cluster/douyin-cluster.yaml


# 创建命名空间
kubectl create namespace douyin-comment

kind create cluster --config deployment/cluster/douyin-cluster.yaml

goctl kube deploy -replicas 2 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name comment-rpc-svc -namespace douyin-comment -i
mage douyin-comment-rpc:v1 -o douyin-comment-rpc.yaml -port 8081 --serviceAccount find-endpoints

goctl kube deploy -nodePort 32010 -replicas 2 -requestCpu 200 -requestMem 50 -limitCpu 300 -limitMem 100 -name douyin-api-svc -namespace do
uyin-comment -image douyin-comment-api:v1 -o douyin-comment-api.yaml -port 8888 --serviceAccount find-endpoints


kubectl apply -f .\douyin-comment-rpc.yaml
kubectl apply -f .\douyin-comment-api.yaml

kind load docker-image douyin-comment-rpc:v1 --name douyin
kind load docker-image douyin-comment-api:v1 --name douyin


