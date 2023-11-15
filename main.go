package main

import (
	"web_service_gin/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	handlers.InitializeHandler()
	router := gin.Default()
	router.GET("/albums", handlers.GetAlbums)
	router.POST("/albums", handlers.PostAlbums)
	router.GET("/albums/:id", handlers.GetAlbumsId)
	router.DELETE("/albums/:id", handlers.DeleteAlbum)
	router.PUT("/albums/:id", handlers.UpdateAlbum)
	router.Run("localhost:8080")

}
