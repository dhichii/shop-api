package main

import (
	"shop-api/src/config/database"
	"shop-api/src/config/env"
	"shop-api/src/router"

	"github.com/joho/godotenv"
)

func main() {
	// load env file
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	// migrate gorm
	database.MigrateMySQL()

	app := router.InitServer()
	app.Listen(":" + env.GetServerEnv())
}
