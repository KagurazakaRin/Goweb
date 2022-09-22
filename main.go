package main

import (
	"goweb/database"
	"goweb/routes"
)

// todo init mysql

func main() {

	database.Connect()

	r := routes.SetRoutes()

	err := r.Run()
	if err != nil {
		panic(err)
	}
}
