package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	dbUri := "mongodb://root:12345@localhost:27017"

	if err := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI(dbUri)); err != nil {
		log.Fatal("Failed to connect to DB. DB string:", dbUri)
	}
}

func main() {
	app := fiber.New()

	configureRoutes(app)

	app.Listen(":3000")
}
