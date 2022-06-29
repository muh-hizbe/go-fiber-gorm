package main

import (
	"go-fiber-gorm/database"
	"go-fiber-gorm/database/migration"
	"go-fiber-gorm/route"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//	INITIAL DATABASE
	database.DatabaseInit()
	//	RUN MIGRATION
	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	errListen := app.Listen(":" + port)
	if errListen != nil {
		log.Println("Fail to listen go fiber server")
		os.Exit(1)
	}
}
