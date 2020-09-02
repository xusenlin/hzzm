package controller

import (
	"github.com/gin-gonic/gin"
	"hzzm/models"
	"net/http"
)

func AuthorDetails(c *gin.Context) {

	name := c.Query("name")

	var author models.Author

	err := author.Seek(name)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": false,
			"data":   "",
			"msg":    err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": true,
		"data":   author,
		"msg":    "成功",
	})

}