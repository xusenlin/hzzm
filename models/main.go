package models

import (
	"gorm.io/gorm"
	"hzzm/db"
	"hzzm/helper"
)

type PageResult struct {
	List      interface{} `json:"list"`
	Total     int64       `json:"total"`
	PageNum   int         `json:"pageNum"`
	PageSize  int         `json:"pageSize"`
	TotalPage int         `json:"totalPage"`
}

func TableMigrate() error {
	return db.Conn.AutoMigrate(&TangShi{}, &SongCi{}, &SongShi{}, &YuanQu{}, &SiShuWuJing{}, &LunYu{}, &YouMengYing{})
}

func getPaginateData(t interface{}, pageNum int, pageSize int, maps,or interface{}, f func(*gorm.DB) error) (*PageResult, error) {
	var result PageResult

	result.PageSize = pageSize
	result.PageNum = pageNum

	query := db.Conn.Model(t).Where(maps).Or(or).Order("star desc")
	query.Count(&result.Total)

	result.TotalPage = helper.ComputeTotalPage(result.Total, int64(pageSize))
	err := f(query.Offset((pageNum - 1) * pageSize).Limit(pageSize))

	if err != nil && err != gorm.ErrRecordNotFound {
		return &result, err
	}

	return &result, nil

}
