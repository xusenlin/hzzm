package controller

import (
	"github.com/gin-gonic/gin"
	"hzzm/models"
	"net/http"
	"strconv"
)

func Paginate(c *gin.Context)  {

	var condition = make(map[string]interface{})
	pageNum, err := strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		pageNum = 1
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil {
		pageSize = 10
	}

	result, err := models.List(c.Query("table"),pageNum, pageSize, condition)

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
		"data":   result,
		"msg":    "查询成功",
	})

}
