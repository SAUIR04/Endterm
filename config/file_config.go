package config

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var (
	MinioClient *minio.Client
	BucketName  = "mybucket"
)

func InitMinioClient() {
	var err error

	MinioClient, err = minio.New("minio:9000", &minio.Options{
		Creds:  credentials.NewStaticV4("admin", "admin12345", ""),
		Secure: false,
	})
	if err != nil {
		log.Fatalln("Error initializing MinIO client:", err)
	}

	ctx := context.Background()

	exists, err := MinioClient.BucketExists(ctx, BucketName)
	if err != nil {
		log.Fatalln("Error checking bucket existence:", err)
	}

	if !exists {
		err = MinioClient.MakeBucket(ctx, BucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalln("Error creating bucket:", err)
		}
		log.Println("Bucket created:", BucketName)
	} else {
		log.Println("Bucket already exists:", BucketName)
	}
}
