package controller

import (
	"encoding/json"
	"os"
	"Endterm/service"
)

type Controller struct {
	service *service.FileService
}

func NewController(s *service.FileService) *Controller {
	return &Controller{service: s}
}

func (c *Controller) UploadTXT() error {
	if err := os.WriteFile("sample.txt", []byte("Это тестовый текстовый файл.\nЗагружен в MinIO/S3 бакет.\nТип: TXT"), 0644); err != nil {
		return err
	}
	defer os.Remove("sample.txt")
	return c.service.UploadFile("sample.txt", "files/text/sample.txt")
}

func (c *Controller) UploadJSON() error {
	data, err := json.MarshalIndent(map[string]interface{}{
		"name": "MinIO Upload Service",
		"version": "1.0.0",
		"tags": []string{"minio", "s3", "storage", "go"},
	}, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile("sample.json", data, 0644); err != nil {
		return err
	}
	defer os.Remove("sample.json")
	return c.service.UploadFile("sample.json", "files/json/sample.json")
}

func (c *Controller) UploadPNG() error {
	png := []byte{0x89,0x50,0x4E,0x47,0x0D,0x0A,0x1A,0x0A,0x00,0x00,0x00,0x0D,0x49,0x48,0x44,0x52,0x00,0x00,0x00,0x01,0x00,0x00,0x00,0x01,0x08,0x06,0x00,0x00,0x00,0x1F,0x15,0xC4,0x89,0x00,0x00,0x00,0x0A,0x49,0x44,0x41,0x54,0x78,0x9C,0x63,0x00,0x01,0x00,0x00,0x05,0x00,0x01,0x0D,0x0A,0x2D,0xB4,0x00,0x00,0x00,0x00,0x49,0x45,0x4E,0x44,0xAE,0x42,0x60,0x82}
	if err := os.WriteFile("sample.png", png, 0644); err != nil {
		return err
	}
	defer os.Remove("sample.png")
	return c.service.UploadFile("sample.png", "files/images/sample.png")
}
