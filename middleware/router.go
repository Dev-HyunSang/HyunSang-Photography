package middleware

import (
	"github.com/dev-hyunsang/photo.parkhyunsang.com-Back-End/functions"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/markbates/pkger"
)

func New(app *fiber.App) {
	app.Use("/assets", filesystem.New(filesystem.Config{
		Root: pkger.Dir("/assets"),
	}))

	app.Use(cors.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET, POST, PUT, DELETE",
	}))

	app.Static("/", "./public")

	api := app.Group("/api")
	api.Get("/", functions.HelloHandler)
	api.Post("/", functions.HelloHandler)
}
