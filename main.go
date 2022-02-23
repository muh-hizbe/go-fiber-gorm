package main

import (
	"github.com/gofiber/fiber/v2"
	"go-fiber-gorm/database"
	"go-fiber-gorm/database/migration"
	"go-fiber-gorm/route"
	"log"
	"os"
)

func main() {
	//	INITIAL DATABASE
	database.DatabaseInit()
	//	RUN MIGRATION
	migration.RunMigration()

	app := fiber.New()

	// INITIAL ROUTE
	route.RouteInit(app)

	errListen := app.Listen(":8080")
	if errListen != nil {
		log.Println("Fail to listen go fiber server")
		os.Exit(1)
	}
}
