package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	dto "torrent-scraper/dto/result"
	"torrent-scraper/models"
	"torrent-scraper/pkg"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/schollz/progressbar/v3"
)

// GetBitsearch godoc
// @Summary Get list bitsearch
// @Description Get list bitsearch
// @Tags bitsearch
// @Param startpage query int false "Start page" Default(1)
// @Param endpage query int false "End page" Default(2)
// @Param keyword query string true "Search torrent list by keyword" Default(adobe)
// @Accept json
// @Produce json
// @Success 200 {object} dto.SuccessResult
// @Failure 500 {object} dto.ErrorResult
// @Router /bitsearch [get]
func BitSearch(ctx *fiber.Ctx) error {
	startTime := time.Now()

	// QUERY PARAMETERS
	// keyword required
	// startpage optional (default = 1)
	// endpage optional (default = 2)

	qkeyword := ctx.Query("keyword")
	qstartPage := ctx.Query("startpage")
	qendPage := ctx.Query("endpage")

	if qkeyword == "" {
		return ctx.JSON(dto.ErrorResult{Status: http.StatusBadRequest, Message: "field keyword is required!"})
	}

	var startPage int
	if qstartPage == "" {
		startPage = 1
	} else {
		startPage, _ = strconv.Atoi(qstartPage)
	}

	var endPage int
	if qendPage == "" {
		endPage = 2
	} else {
		endPage, _ = strconv.Atoi(qendPage)
	}

	c := colly.NewCollector()

	var posts []models.BitSearch

	c.OnHTML(".search-result", func(e *colly.HTMLElement) {
		post := models.BitSearch{}

		post.Date = e.ChildText(".stats > div:nth-child(5)")
		post.Title = e.ChildText(".title > a")
		post.Size = e.ChildText(".stats > div:nth-child(2)")
		post.Seeder = e.ChildText(".stats > div:nth-child(3)")
		post.Leecher = e.ChildText(".stats > div:nth-child(4)")
		post.Downloader = e.ChildText(".stats > div:nth-child(1)")

		e.ForEach(".title > a", func(_ int, link *colly.HTMLElement) {
			post.Link = "https://bitsearch.to" + link.Attr("href")
		})

		e.ForEach(".links > a.dl-torrent", func(_ int, link *colly.HTMLElement) {
			post.TorrentLink = link.Attr("href")
		})

		e.ForEach(".links > a.dl-magnet", func(_ int, link *colly.HTMLElement) {
			post.MagnetURL = link.Attr("href")
		})

		posts = append(posts, post)

	})

	bar := progressbar.Default(int64(endPage))

	for page := startPage; page <= endPage; page++ {
		url := fmt.Sprintf("https://bitsearch.to/search?q=%s&page=%d", qkeyword, page)
		c.Visit(url)
		bar.Add(page)
	}

	bar.Finish()

	headers := []string{"Date", "Title", "Size", "Seeder", "Leecher", "Downloader", "Link", "Torrent Link", "Magnet URL"}
	pkg.SaveToCSV(posts, "bitsearch_scrape.csv", headers)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Waktu eksekusi: %s\n", elapsedTime)

	return ctx.JSON(dto.SuccessResult{Status: http.StatusOK, Data: posts})
}
