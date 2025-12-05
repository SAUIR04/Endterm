package service

import (
	"context"
	"fmt"
	"mime"
	"path/filepath"
	"time"

	"Endterm/config"
	"Endterm/models"

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

func (s *FileService) UploadFile(filePath string, objectName string) (*models.File, error) {
	contentType := mime.TypeByExtension(filepath.Ext(filePath))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	info, err := s.client.FPutObject(s.ctx, s.bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		return nil, fmt.Errorf("ошибка при загрузке файла в MinIO: %w", err)
	}

	f := &models.File{
		ID:           objectName,
		Name:         objectName,
		OriginalName: filepath.Base(filePath),
		Size:         info.Size,
		ContentType:  contentType,
		UploadedAt:   time.Now(),
	}

	return f, nil
}
