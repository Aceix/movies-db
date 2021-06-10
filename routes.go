package main

import (
	"github.com/gofiber/fiber/v2"
	controllers "theaceix.site/movies-db/controllers"
)

func configureRoutes(app *fiber.App) {
	app.Get("/", controllers.GetHome)
}
