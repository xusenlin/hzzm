package models

//幽梦影
import (
	"gorm.io/gorm"
	"hzzm/db"
	"time"
)

type YouMengYing struct {
	ID        uint       `gorm:"primary_key"json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"json:"deleted_at"`
	Star      int        `json:"star"`
	Content   string     `json:"content"`
	Comment   string     `json:"comment"`
}

func (YouMengYing) TableName() string {
	return "you_meng_ying"
}
func (t *YouMengYing) List(pageNum int, pageSize int, maps interface{}) (*PageResult, error) {

	var r []YouMengYing

	result, err := getPaginateData(t, pageNum, pageSize, maps, func(q *gorm.DB) error { return q.Find(&r).Error })

	if err != nil {
		return result, err
	}
	result.List = r

	return result, nil
}
func (t *YouMengYing) Favour(id string) error {
	return db.Conn.Model(&t).Where("id = ?", id).
		UpdateColumn("star", gorm.Expr("star + ?", 1)).Error
}