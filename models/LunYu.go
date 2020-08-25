package models

//四书五经

import (
	"gorm.io/gorm"
	"time"
)

type LunYu struct {
	ID         uint       `gorm:"primary_key"json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index"json:"deleted_at"`
	Star       int        `json:"star"`
	Chapter    string     `json:"chapter"`
	Paragraphs string     `json:"paragraphs"`
}

func (LunYu) TableName() string {
	return "lun_yu"
}

func (t *LunYu) List(pageNum int, pageSize int, maps interface{}) (*PageResult, error) {

	var r []LunYu

	result, err := getPaginateData(t, pageNum, pageSize, maps, func(q *gorm.DB) error { return q.Find(&r).Error })

	if err != nil {
		return result, err
	}
	result.List = r

	return result, nil
}
