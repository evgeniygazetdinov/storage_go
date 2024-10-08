package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var path, err = os.Getwd()

func getAlbums(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, models.Albums)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Posts",
	})
}

func getAlbums2(c *gin.Context) {

	c.HTML(http.StatusOK, "main.tmpl", gin.H{
		"title": "Posts",
	})
	// c.IndentedJSON(http.StatusOK, models.GetUser())

}
func getAlbums3(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, "{'HELL': 'YEAH'}")

}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("temp/*")
	router.GET("/albums", getAlbums)
	router.GET("/albums2", getAlbums2)
	router.GET("/forCheck", getAlbums3)
	router.Run("localhost:8085")
}
