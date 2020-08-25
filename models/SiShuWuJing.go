package models

//四书五经

import (
	"gorm.io/gorm"
	"time"
)

type SiShuWuJing struct {
	ID         uint       `gorm:"primary_key"json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index"json:"deleted_at"`
	Star       int        `json:"star"`
	Chapter    string     `json:"chapter"`
	Paragraphs string     `json:"paragraphs"`
}

func (SiShuWuJing) TableName() string {
	return "si_shu_wu_jing"
}
func (t *SiShuWuJing) List(pageNum int, pageSize int, maps interface{}) (*PageResult, error) {

	var r []SiShuWuJing

	result, err := getPaginateData(t, pageNum, pageSize, maps, func(q *gorm.DB) error { return q.Find(&r).Error })

	if err != nil {
		return result, err
	}
	result.List = r

	return result, nil
}
