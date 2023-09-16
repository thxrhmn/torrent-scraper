package main

import (
	"fmt"
	"net/http"
	"os"
	dto "torrent-scraper/dto/result"
	"torrent-scraper/handlers"

	_ "torrent-scraper/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

// @title           Torrent Scraper API
// @version         1.0
// @description     This is a Torrent Scraper API server.
// @contact.name    API Support
// @contact.url     http://github.com/thxrhmn/torrent-scraper
// @host            localhost:3000
// @BasePath        /api/v1

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/", func(c *fiber.Ctx) error {
		data := "hello world"
		return c.JSON(dto.SuccessResult{Status: http.StatusOK, Data: data})
	})

	v1 := app.Group("/api/v1")
	v1.Get("/btdig", handlers.BtDIg)
	v1.Get("/bitsearch", handlers.BitSearch)

	fmt.Println("Server listening on port:", port)

	app.Listen(":" + port)

}
