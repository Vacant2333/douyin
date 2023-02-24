# Deploy this project

## 1. Install dependencies
If you use Windows, we recommend that you install dependent software in the virtual machine.
- [Docker](https://www.docker.com/)
- [Kind](https://kind.sigs.k8s.io/)
- [Kubectl](https://kubernetes.io/docs/tasks/tools/)
- [Helm](https://helm.sh/docs/intro/install/)
- [Make](https://www.gnu.org/software/make/#download)

## 2. Install the cluster and components
Run this commands in root direction.
#### Install Cluster
```bash
make install-cluster
```

#### Install components which you need
```bash
make install-minio install-etcd install-kafka install-redis
```

## Install the microservices
```bash
make install-minio-client install-user install-comment install-mq install-favorite install-video install-message install-follow install-gateway
```

## Optional components
- [Kubernetes-Dashboard](https://github.com/Vacant2333/douyin/blob/main/deployment/dashboard/README.md)
