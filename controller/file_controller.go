package controller

import (
	"encoding/json"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"Endterm/models"
	"Endterm/service"
)

type Controller struct {
	service *service.FileService
}

func NewController(s *service.FileService) *Controller {
	return &Controller{service: s}
}

func (c *Controller) UploadHandler(w http.ResponseWriter, r *http.Request) {
	r.Body = http.MaxBytesReader(w, r.Body, 50<<20)
	if err := r.ParseMultipartForm(50 << 20); err != nil {
		http.Error(w, "ошибка парсинга multipart form: "+err.Error(), http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "поле file обязательно: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	uploaded, err := c.saveAndUpload(file, header)
	if err != nil {
		log.Println("upload error:", err)
		http.Error(w, "ошибка загрузки: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(uploaded)
}

func (c *Controller) saveAndUpload(file multipart.File, header *multipart.FileHeader) (*models.File, error) {

	tempFile, err := os.CreateTemp("", "upload-*"+filepath.Ext(header.Filename))
	if err != nil {
		return nil, err
	}

	defer func() {
		tempFile.Close()
		os.Remove(tempFile.Name())
	}()

	if _, err := io.Copy(tempFile, file); err != nil {
		return nil, err
	}

	objectName := filepath.Join("", time.Now().Format("20060102-150405")+"-"+header.Filename)

	return c.service.UploadFile(tempFile.Name(), objectName)
}
