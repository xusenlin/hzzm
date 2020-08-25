package helper

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func InStrArr(str string, arr []string) bool {
	for _, item := range arr {
		if item == str {
			return true
		}
	}
	return false
}

func ComputeTotalPage(total int64, pageSize int64, ) int {
	totalPage := total / pageSize
	if total%pageSize != 0 {
		totalPage++
	}
	return int(totalPage)
}

func ParsePageParams(c *gin.Context) (pageNum, pageSize int) {
	var err error
	pageNum, err = strconv.Atoi(c.Query("pageNum"))
	if err != nil {
		pageNum = 1
	}
	pageSize, err = strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize > 500 {
		pageSize = 10
	}
	return
}
