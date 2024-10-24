package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ConvertQueryDataToValues(c *gin.Context, stringKey string) int {
	data := c.Request.URL.Query().Get(stringKey)
	keyValue, _ := strconv.Atoi(data)
	return keyValue
}
