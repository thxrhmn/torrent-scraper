package main

import (
	dto "detik-scraper-gocolly/dto/result"
	"detik-scraper-gocolly/handlers"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		data := "hello world"
		return c.JSON(dto.SuccessResult{Status: http.StatusOK, Data: data})
	})

	v1 := app.Group("/api/v1")
	v1.Get("/btdig", handlers.BtDIg)

	app.Listen(":" + port)
}
