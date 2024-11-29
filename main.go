package main

import (
	"go-agreenery/database"
	"go-agreenery/routes"
	"log"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
	loadEnv()

	db, _ := database.InitDB()

	database.MigrateDB(db)

	e := echo.New()

	routes.InitRoutes(e, db)

	e.Logger.Fatal(e.Start(":1323"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("failed to load env")
	}
}
