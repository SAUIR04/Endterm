package service

import (
	"context"
	"fmt"
	"log"
	"mime"
	"path/filepath"

	"Endterm/config"

	"github.com/minio/minio-go/v7"
)

type FileService struct {
	client     *minio.Client
	bucketName string
	ctx        context.Context
}

func NewFileService() *FileService {
	return &FileService{
		client:     config.MinioClient,
		bucketName: config.BucketName,
		ctx:        context.Background(),
	}
}

func (s *FileService) UploadFile(filePath string, objectName string) error {
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	info, err := s.client.FPutObject(s.ctx, s.bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return fmt.Errorf("ошибка при загрузке файла: %w", err)
	}

	log.Printf("✓ Файл '%s' загружен. Размер: %d байт", objectName, info.Size)
	return nil
}
