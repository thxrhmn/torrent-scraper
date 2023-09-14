package main

import (
	"net/http"
	"os"
	dto "torrent-scraper/dto/result"
	"torrent-scraper/handlers"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
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
