package main

import (
	"log"
	"net/http"
	"time"

	"Endterm/config"
	"Endterm/controller"
	"Endterm/service"
)

func main() {

	config.InitMinioClient()

	fs := service.NewFileService()
	ctrl := controller.NewController(fs)

	http.HandleFunc("/upload", ctrl.UploadHandler)

	addr := ":8080"
	srv := &http.Server{
		Addr:         addr,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}

	log.Println("Server started at", addr)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("server error:", err)
	}
}
