package logic

import (
	"douyin/pkg/logger"
	"github.com/minio/minio-go/v7"
)

import (
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func MakeMinIOClient() *minio.Client {
	var endpoint = "minio.minio.svc.cluster.local:9000"
	var accessKeyID = "douyin"
	var secretAccessKey = "douyin_pass"

	// Create a client, all operations must use it
	client, err := minio.New(endpoint, &minio.Options{
		Creds: credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
	})
	if err != nil {
		logger.Fatalf("Make MinIO client error:", err)
		panic(err)
	}
	return client
}
