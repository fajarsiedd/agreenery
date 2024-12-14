package main

import (
	"go-agreenery/database"
	"go-agreenery/helpers"
	"go-agreenery/routes"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
)

func main() {
	loadEnv()

	db, _ := database.InitDB()

	database.MigrateDB(db)

	e := echo.New()

	routes.InitRoutes(e, db)

	c := cron.New(cron.WithLocation(time.FixedZone("Asia/Bangkok", 7*60*60)))

	c.AddFunc("0 6 * * *", func() {
		helpers.SendWateringScheduleNotifications(db)
	})

	c.AddFunc("0 12 * * *", func() {
		helpers.SendAdminNotifications(db)
	})

	c.Start()

	e.Logger.Fatal(e.Start(":1323"))
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		panic("failed to load env")
	}
}
