package models

//唐诗宋词元曲

import (
	"gorm.io/gorm"
	"hzzm/db"
	"time"
)

type BasePoetry struct {
	ID        uint       `gorm:"primary_key"json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"json:"deleted_at"`
	Title     string     `json:"title"`
	Star      int        `json:"star"`
	Author    string     `json:"author"`
}

type TangShi struct {
	//唐诗
	BasePoetry
	Paragraphs string `json:"paragraphs"`
}

func (TangShi) TableName() string {
	return "tang_shi"
}

type SongShi struct {
	//诗
	BasePoetry
	Paragraphs string `json:"paragraphs"`
}

func (SongShi) TableName() string {
	return "song_shi"
}

type SongCi struct {
	//宋词
	BasePoetry
	Paragraphs string `json:"paragraphs"`
}

func (SongCi) TableName() string {
	return "song_ci"
}

type YuanQu struct {
	//元曲
	BasePoetry
	Paragraphs string `json:"paragraphs"`
}

func (YuanQu) TableName() string {
	return "yuan_qu"
}

func (t *TangShi) List(pageNum int, pageSize int, keyword string) (*PageResult, error) {

	var r []TangShi
	var maps,or map[string]interface{}

	if keyword != ""{
		maps = map[string]interface{}{"title": keyword}
		or = map[string]interface{}{"author": keyword}
	}

	result, err := getPaginateData(t, pageNum, pageSize, maps , or, func(q *gorm.DB) error { return q.Find(&r).Error })

	if err != nil {
		return result, err
	}
	result.List = r

	return result, nil
}
func (t *SongShi) List(pageNum int, pageSize int, keyword string) (*PageResult, error) {

	var r []SongShi

	var maps,or map[string]interface{}

	if keyword != ""{
		maps = map[string]interface{}{"title": keyword}
		or = map[string]interface{}{"author": keyword}
	}

	result, err := getPaginateData(t, pageNum, pageSize, maps,or, func(q *gorm.DB) error { return q.Find(&r).Error })

	if err != nil {
		return result, err
	}
	result.List = r

	return result, nil
}
func (t *SongCi) List(pageNum int, pageSize int, keyword string) (*PageResult, error) {

	var r []SongCi
	var maps,or map[string]interface{}

	if keyword != ""{
		maps = map[string]interface{}{"title": keyword}
		or = map[string]interface{}{"author": keyword}
	}

	result, err := getPaginateData(t, pageNum, pageSize, maps,or, func(q *gorm.DB) error { return q.Find(&r).Error })

	if err != nil {
		return result, err
	}
	result.List = r

	return result, nil
}
func (t *YuanQu) List(pageNum int, pageSize int, keyword string) (*PageResult, error) {

	var r []YuanQu
	var maps,or map[string]interface{}

	if keyword != ""{
		maps = map[string]interface{}{"title": keyword}
		or = map[string]interface{}{"author": keyword}
	}

	result, err := getPaginateData(t, pageNum, pageSize, maps,or, func(q *gorm.DB) error { return q.Find(&r).Error })

	if err != nil {
		return result, err
	}
	result.List = r

	return result, nil
}

func (t *TangShi) Favour(id string) error {
	return db.Conn.Model(&t).Where("id = ?", id).
		UpdateColumn("star", gorm.Expr("star + ?", 1)).Error
}
func (t *SongShi) Favour(id string) error {
	return db.Conn.Model(&t).Where("id = ?", id).
		UpdateColumn("star", gorm.Expr("star + ?", 1)).Error
}
func (t *SongCi) Favour(id string) error {
	return db.Conn.Model(&t).Where("id = ?", id).
		UpdateColumn("star", gorm.Expr("star + ?", 1)).Error
}
func (t *YuanQu) Favour(id string) error {
	return db.Conn.Model(&t).Where("id = ?", id).
		UpdateColumn("star", gorm.Expr("star + ?", 1)).Error
}
