package logic

import (
	"bytes"
	"context"
	"douyin/pkg/logger"
	"fmt"
	"github.com/minio/minio-go/v7"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
	"time"
)

import (
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func makeMinIOClient() *minio.Client {
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

func getVideoFrame(data []byte, frame int) (*bytes.Reader, error) {
	tmp, _ := os.Create("tmp.mp4")
	tmp.Write(data)
	tmp.Close()

	pngBuffer := bytes.NewBuffer(nil)
	err := ffmpeg.Input("tmp.mp4").Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frame)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(pngBuffer).Run()
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(pngBuffer.Bytes()), nil
}

func uploadFile(client *minio.Client, reader *bytes.Reader, fileName string, bucket string, contentType string) (string, error) {
	_, err := client.PutObject(context.Background(),
		bucket, fileName, reader, reader.Size(), minio.PutObjectOptions{})
	//minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		logger.Fatalf("Fail to upload file, name: %v, err: %v", fileName, err)
		return "", err
	}

	fileUrl, err := client.PresignedGetObject(context.Background(), bucket, fileName, time.Second*60*60*24*7, nil)
	if err != nil {
		logger.Fatalf("Fail to get file url, fileName: %v, err: %v", fileName, err)
		return "", err
	}

	logger.InfoF("Success to upload object to minio, fileName: %v, url: %v", fileName, fileUrl.String())
	return fileUrl.String(), nil
}
