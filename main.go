package main

import (
	"log"
	"Endterm/config"
	"Endterm/controller"
	"Endterm/service"
)

func main() {
	log.Println("Инициализация MinIO...")
	config.InitMinioClient()

	s := service.NewFileService()
	c := controller.NewController(s)

	if err := c.UploadTXT(); err != nil {
		log.Fatalf("TXT error: %v", err)
	}

	if err := c.UploadJSON(); err != nil {
		log.Fatalf("JSON error: %v", err)
	}

	if err := c.UploadPNG(); err != nil {
		log.Fatalf("PNG error: %v", err)
	}

	log.Println("✓ Все файлы успешно загружены!")
}
