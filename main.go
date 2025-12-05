//package Endterm
//
//import (
//	"github.com/gin-gonic/gin"
//	"minIO/config"
//	"minIO/controllers"
//	"minIO/services"
//)
//
//func main() {
//	config.InitMinioClient()
//	config.InitPostgres()
//
//	minioService := services.NewService(config.MinioClient, config.DB)
//	controller := controllers.NewController(minioService)
//
//	r := gin.Default()
//	controller.RegisterRoutes(r)
//	r.Run(":8080")
//}
