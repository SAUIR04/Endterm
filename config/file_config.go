package config

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client
var BucketName = "mybucket"
var ctx = context.Background()

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

func InitMinioClient() {
	endpoint := getEnv("MINIO_ENDPOINT", "localhost:9000")
	accessKey := getEnv("MINIO_ACCESS_KEY", "admin")
	secretKey := getEnv("MINIO_SECRET_KEY", "admin12345")
	useSSL := getEnv("MINIO_SECURE", "false") == "true"
	BucketName = getEnv("MINIO_BUCKET", "mybucket")

	log.Printf("Подключение к MinIO: %s\n", endpoint)

	var err error
	for i := 0; i < 10; i++ {
		MinioClient, err = minio.New(endpoint, &minio.Options{
			Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
			Secure: useSSL,
		})
		if err == nil {
			break
		}
		log.Printf("Попытка подключения %d/10...\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalln("Ошибка подключения к MinIO:", err)
	}

	log.Println("✓ Подключение к MinIO установлено")

	exists, err := MinioClient.BucketExists(ctx, BucketName)
	if err != nil {
		log.Fatalln("Ошибка проверки бакета:", err)
	}

	if !exists {
		err = MinioClient.MakeBucket(ctx, BucketName, minio.MakeBucketOptions{})
		if err != nil {
			log.Fatalln("Ошибка создания бакета:", err)
		}
		log.Printf("✓ Бакет '%s' создан\n", BucketName)
	} else {
		log.Printf("✓ Бакет '%s' уже существует\n", BucketName)
	}
}
