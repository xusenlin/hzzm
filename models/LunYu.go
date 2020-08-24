package models

//四书五经

import (
	"time"
)

type LunYu struct {
	ID         uint       `gorm:"primary_key"json:"id"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
	DeletedAt  *time.Time `sql:"index"json:"deleted_at"`
	Chapter    string     `json:"chapter"`
	Paragraphs string     `json:"paragraphs"`
}
func (LunYu) TableName() string {
	return "lun_yu"
}