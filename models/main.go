package models

import (
	"errors"
	"gorm.io/gorm"
	"hzzm/db"
	"hzzm/helper"
)

type PageResult struct {
	List      interface{}
	Total     int64
	PageNum   int
	PageSize  int
	TotalPage int
}

var TableAlias = map[string]string{
	"ts":   "tang_shi",
	"sc":   "song_ci",
	"yq":   "yuan_qu",
	"ly":   "lun_yu",
	"sswj": "si_shu_wu_jing",
	"ymy":  "you_meng_ying",
}

func List(tableAlias string, pageNum int, pageSize int, maps interface{}) (PageResult, error) {
	var result PageResult


	var err error

	result.List, result.Total, err = getData(tableAlias, pageNum, pageSize, maps)
	if err != nil {
		return result, err
	}

	result.PageSize = pageSize
	result.PageNum = pageNum

	result.TotalPage = helper.ComputeTotalPage(result.Total, int64(pageSize))

	return result, nil
}

func getData(tableAlias string, pageNum int, pageSize int, maps interface{}) (interface{}, int64, error) {

	var total int64
	var result interface{}

	t,ok := TableAlias[tableAlias]
	if !ok {
		return result, 0,errors.New("不存在的表别名")
	}

	query := db.Conn.Table(t).Where(maps).Order("star desc")

	query.Count(&total)
	//err := query.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&result).Error

	query = query.Offset((pageNum - 1) * pageSize).Limit(pageSize)

	var err error
	switch tableAlias {
	case "ts":
		var r []*TangShi
		err = query.Find(&r).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return r, 0, err
		}
		return r, total, nil
	case "sc":
		var r []*SongCi
		err = query.Find(&r).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return r, 0, err
		}
		return r, total, nil
	case "yq":
		var r []*YuanQu
		err = query.Find(&r).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return r, 0, err
		}
		return r, total, nil
	case "ly":
		var r []*LunYu
		err = query.Find(&r).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return r, 0, err
		}
		return r, total, nil
	case "sswj":
		var r []*SiShuWuJing
		err = query.Find(&r).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return r, 0, err
		}
		return r, total, nil
	case "ymy":
		var r []*YouMengYing
		err = query.Find(&r).Error
		if err != nil && err != gorm.ErrRecordNotFound {
			return r, 0, err
		}
		return r, total, nil
	}
	return result, 0,errors.New("不存在的表")
}
