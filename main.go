package main

import (
	"example/storage-go/models"
	"fmt"
	"net/http"
	"strconv"

	helpers "example/storage-go/helpers"

	"github.com/gin-gonic/gin"
)

func getDays(c *gin.Context) {
	// c.IndentedJSON(http.StatusOK, models.Albums)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Posts",
	})
}

func checkDay(c *gin.Context) {
	data := c.Request.URL.Query().Get("id")
	user_id, _ := strconv.Atoi(data)

	user := models.GetUser(user_id)
	fmt.Println(user)
	c.IndentedJSON(http.StatusOK, "{'HELL': 'YEAH'}")

}

func addEvent(c *gin.Context) {
	myEventId := helpers.ConvertQueryDataToValues(c, "id")
	fmt.Println(myEventId)
	c.IndentedJSON(http.StatusOK, "{'HELL':"+string(myEventId)+"}")
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("temp/*")
	router.GET("/", getDays)
	router.GET("/forCheck", checkDay)
	router.POST("/addEvent", addEvent)
	router.Run("localhost:8085")
}
