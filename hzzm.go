package main

import (
	"hzzm/config"
	"hzzm/db"
	"hzzm/models"
	"hzzm/routes"
)

func main() {

	db.InitConn()

	if err := models.TableMigrate(); err != nil {
		panic(err)
	}

	r := routes.InitRouter()

	if err := r.Run(":" + config.HttpPort); err != nil {
		panic(err)
	}
}
