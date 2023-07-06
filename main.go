package main

import (
	"log"
	"soduku-madness-BE/internal/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		AppName:      "Soduku-solver API",
		ServerHeader: "Fiber",
	})

	app.Mount("/api/v1", router.GetAPI())

	log.Fatal(app.Listen(":3000"))
}
