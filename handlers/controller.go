package handlers

import (
	"net/http"
	"web_service_gin/schemas"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	alb := []schemas.Album{}
	result := db.Find(&alb)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
	}
	c.JSON(http.StatusFound, alb)
}

func PostAlbums(c *gin.Context) {
	var alb schemas.Album

	if err := c.BindJSON(&alb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result := db.Create(&alb)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
	}

	c.JSON(http.StatusCreated, alb.ID)

}

func GetAlbumsId(c *gin.Context) {
	var alb schemas.Album
	id := c.Param("id")

	result := db.First(&alb, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, result.Error)
	}
	c.JSON(http.StatusFound, alb)

}

func DeleteAlbum(c *gin.Context) {
	var alb schemas.Album
	id := c.Param("id")

	result := db.Delete(&alb, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}

func UpdateAlbum(c *gin.Context) {
	var updtAlb schemas.Album
	id := c.Param("id")

	if err := c.BindJSON(&updtAlb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var alb schemas.Album
	result := db.First(&alb, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	if updtAlb.Artist != "" {
		alb.Artist = updtAlb.Artist
	}

	if updtAlb.Title != "" {
		alb.Title = updtAlb.Title
	}

	if updtAlb.Price > 0 {
		alb.Price = updtAlb.Price
	}
	result = db.Save(&alb)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, updtAlb)
}
