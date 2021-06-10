package controllers

import "github.com/gofiber/fiber/v2"

func GetHome(ctx *fiber.Ctx) error {
	return ctx.SendString("Hello world!")
}
