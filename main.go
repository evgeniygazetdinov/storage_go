package main

import (
	"net/http"

	"example/storage-go/models"

	"github.com/gin-gonic/gin"
)

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Albums)
}

func getAlbums2(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.GetUser())
}

func main() {
	router := gin.Default()
	router.GET("/albums", getAlbums)
	router.GET("/albums2", getAlbums2)
	router.Run("localhost:8080")
}
