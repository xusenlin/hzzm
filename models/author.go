package models

import (
	"gorm.io/gorm"
	"hzzm/db"
	"time"
)

type Author struct {
	ID        uint       `gorm:"primary_key"json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `sql:"index"json:"deleted_at"`
	Name      string     `json:"name"`
	Star      int        `json:"star"`
	Desc      string     `json:"desc"`
	Dynasty   string     `json:"dynasty"`
}
func (t *Author) Favour(id string) error {
	return db.Conn.Model(&t).Where("id = ?", id).
		UpdateColumn("star", gorm.Expr("star + ?", 1)).Error
}