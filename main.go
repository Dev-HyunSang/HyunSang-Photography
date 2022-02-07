/*
Developer. HyunSang Park, parkhyunsang@kakao.com
*/
package main

import (
	"log"

	"github.com/dev-hyunsang/photo.parkhyunsang.com-Back-End/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	middleware.New(app)

	log.Println("Starting Server...")
	if err := app.Listen(":80"); err != nil {
		log.Fatal("Failed Running Server...!")
	}
}
