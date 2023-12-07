package main

import (
	"fmt"
	"net/http"
	"os"
	dto "torrent-scraper/dto/result"
	"torrent-scraper/handlers"

	_ "torrent-scraper/docs"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
)

// @title           Torrent Scraper API
// @version         1.0
// @description     This is a Torrent Scraper API server.
// @contact.name    API Support
// @contact.url     http://github.com/thxrhmn/torrent-scraper
// @host            127.0.0.1:8080
// @BasePath        /api/v1

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	app := fiber.New()

	// Initialize default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
		AllowMethods: "*",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/", func(c *fiber.Ctx) error {
		data := "hello world"
		return c.JSON(dto.SuccessResult{Status: http.StatusOK, Data: data})
	})

	app.Static("/files", "./results")

	v1 := app.Group("/api/v1")
	v1.Get("/btdig", handlers.BtDIg)
	v1.Get("/bitsearch", handlers.BitSearch)
	v1.Get("/files", handlers.GetFiles)

	fmt.Println("Server listening on port:", port)

	app.Listen(":" + port)

}
