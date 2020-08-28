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

func (t *TangShi) List(pageNum int, pageSize int, maps interface{}) (*PageResult, error) {

	var r []TangShi

	result, err := getPaginateData(t, pageNum, pageSize, maps, func(q *gorm.DB) error { return q.Find(&r).Error })

	if err != nil {
		return result, err
	}
	result.List = r

	return result, nil
}
func (t *SongCi) List(pageNum int, pageSize int, maps interface{}) (*PageResult, error) {

	var r []SongCi

	result, err := getPaginateData(t, pageNum, pageSize, maps, func(q *gorm.DB) error { return q.Find(&r).Error })

	if err != nil {
		return result, err
	}
	result.List = r

	return result, nil
}
func (t *YuanQu) List(pageNum int, pageSize int, maps interface{}) (*PageResult, error) {

	var r []YuanQu

	result, err := getPaginateData(t, pageNum, pageSize, maps, func(q *gorm.DB) error { return q.Find(&r).Error })

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
func (t *SongCi) Favour(id string) error {
	return db.Conn.Model(&t).Where("id = ?", id).
		UpdateColumn("star", gorm.Expr("star + ?", 1)).Error
}
func (t *YuanQu) Favour(id string) error {
	return db.Conn.Model(&t).Where("id = ?", id).
		UpdateColumn("star", gorm.Expr("star + ?", 1)).Error
}