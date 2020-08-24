package main

import (
	"hzzm/config"
	"hzzm/db"
	"hzzm/routes"
)

func main() {
	db.InitConn()

	r := routes.InitRouter()

	err := r.Run(":" + config.HttpPort)
	if err != nil {
		panic(err)
	}
}
