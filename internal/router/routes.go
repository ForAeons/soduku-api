package router

import (
	"soduku-madness-BE/internal/handler/health"
	"soduku-madness-BE/internal/handler/soduku"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func GetAPI() *fiber.App {
	app := fiber.New()

	app.Use(cors.New())
	// app.Use(limiter.New())
	app.Use(logger.New())
	app.Use(recover.New())

	app.Get("/health", health.HandleHealth)

	sodukuapi := app.Group("/soduku")
	sodukuapi.Post("/solve", soduku.HandleSolve)
	sodukuapi.Get("/generate", soduku.HandleGeneratePuzzle)

	return app
}
