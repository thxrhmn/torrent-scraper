package handlers

import (
	"net/http"
	dto "torrent-scraper/dto/result"
	"torrent-scraper/pkg"

	"github.com/gofiber/fiber/v2"
)

func GetFiles(ctx *fiber.Ctx) error {
	files := pkg.GetFileResults()

	return ctx.JSON(dto.SuccessResult{Status: http.StatusOK, Data: files})
}
