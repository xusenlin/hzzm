package db


import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Conn *gorm.DB

func InitConn() {
	var err error

	Conn, err = gorm.Open(sqlite.Open("./database.db"), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}