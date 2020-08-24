package models

//幽梦影
import "time"

type YouMengYing struct {
	ID        uint       `gorm:"primary_key"json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"json:"deleted_at"`
	Content   string     `json:"content"`
	Comment   string     `json:"comment"`
}
func (YouMengYing) TableName() string {
	return "you_meng_ying"
}