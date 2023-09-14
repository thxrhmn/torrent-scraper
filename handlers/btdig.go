package handlers

import (
	dto "detik-scraper-gocolly/dto/result"
	"detik-scraper-gocolly/models"
	"detik-scraper-gocolly/pkg"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/schollz/progressbar/v3"
)

func BtDIg(ctx *fiber.Ctx) error {
	startTime := time.Now()

	queryStartPage := ctx.Query("startpage")
	queryEndPage := ctx.Query("endpage")
	keyword := ctx.Query("keyword")
	order := ctx.Query("order")

	// order by relevance | age | size | files

	// 1 = relevance
	// 2 = age
	// 3 = size
	// 4 = files

	var startPage int
	if queryStartPage != "" {
		startPage, _ = strconv.Atoi(queryStartPage)
	} else {
		startPage = 1
	}

	var endPage int
	if queryEndPage != "" {
		endPage, _ = strconv.Atoi(queryEndPage)
	} else {
		endPage = 2
	}

	var orderby int
	switch order {
	case "relevance":
		orderby = 1
	case "age":
		orderby = 2
	case "size":
		orderby = 3
	case "files":
		orderby = 4
	default:
		orderby = 1
	}

	if keyword == "" {
		return ctx.Status(http.StatusBadRequest).JSON(dto.ErrorResult{Status: http.StatusBadRequest, Message: "please fill the keyword"})
	}

	c := colly.NewCollector()

	var posts []models.Btdig

	c.OnHTML(".one_result", func(e *colly.HTMLElement) {
		post := models.Btdig{}
		post.Date = e.ChildText(".torrent_age")
		post.Title = e.ChildText(".torrent_name > a")
		post.Size = e.ChildText(".torrent_size")
		post.Files = e.ChildText(".torrent_files") + " Files"

		e.ForEach(".torrent_name > a", func(_ int, link *colly.HTMLElement) {
			post.Link = link.Attr("href")
		})

		e.ForEach(".torrent_magnet .fa-magnet > a", func(_ int, magnet *colly.HTMLElement) {
			post.MagnetURL = magnet.Attr("href")
		})

		posts = append(posts, post)
	})

	bar := progressbar.Default(int64(endPage))

	for page := startPage; page <= endPage; page++ {
		url := fmt.Sprintf("https://btdig.com/search?q=%s&p=%d&order=%d", keyword, page, orderby)
		fmt.Println("scraping....", url)
		err := c.Visit(url)
		if err != nil {
			log.Fatal(err)
		}

		bar.Add(page)
	}

	bar.Finish()

	headers := []string{"Date", "Title", "Files", "Size", "Link", "Magnet URL"}
	pkg.SaveToCSV(posts, "btdig_scrape.csv", headers)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Waktu eksekusi: %s\n", elapsedTime)

	return ctx.JSON(dto.SuccessResult{Status: http.StatusOK, Data: posts})
}
