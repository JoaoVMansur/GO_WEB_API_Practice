package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	var albums []Album
	if db == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database connection is not established"})
		return
	}

	rows, err := db.Query("SELECT * FROM Album")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}
	defer rows.Close()

	for rows.Next() {
		var alb Album
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		albums = append(albums, alb)

	}
	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, albums)

}

func postAlbums(c *gin.Context) {
	var alb Album

	if err := c.BindJSON(&alb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	result, err := db.Exec("INSERT INTO album (title,artist,price) VALUES (?,?,?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	id, err := result.LastInsertId()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, id)
}

func getAlbumsId(c *gin.Context) {
	var alb Album
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		fmt.Println("Error:", err)
	}

	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	c.JSON(http.StatusOK, alb)

}
func deleteAlbum(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	result, err := db.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	rowsAffected, err := result.RowsAffected()
	if rowsAffected == 0 || err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Album not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Album deleted successfully"})
}

func updateAlbum(c *gin.Context) {
	var updtAlb Album
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.BindJSON(&updtAlb); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result, err := db.Exec("UPDATE album SET Title = ?, Artist = ?, Price = ? WHERE id = ?", updtAlb.Title, updtAlb.Artist, updtAlb.Price, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, result)
}
