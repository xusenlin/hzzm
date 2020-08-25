package controller

import (
	"github.com/gin-gonic/gin"
	"hzzm/helper"
	"hzzm/models"
	"net/http"
)

func TangShiPaginate(c *gin.Context) {

	pageNum, pageSize := helper.ParsePageParams(c)

	result, err := new(models.TangShi).List(pageNum, pageSize, map[string]string{})

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
