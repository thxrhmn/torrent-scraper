package handlers

import (
	"fmt"
	"time"
	"torrent-scraper/models"
	"torrent-scraper/pkg"

	"github.com/gocolly/colly/v2"
	"github.com/schollz/progressbar/v3"
)

func BitSearch(keyword string, totalPage int) {
	startTime := time.Now()

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

	bar := progressbar.Default(int64(totalPage))

	for page := 0; page < totalPage; page++ {
		url := fmt.Sprintf("https://bitsearch.to/search?q=%s&page=%d", keyword, page)
		c.Visit(url)
		bar.Add(page)
	}

	bar.Finish()

	headers := []string{"Date", "Title", "Size", "Seeder", "Leecher", "Downloader", "Link", "Torrent Link", "Magnet URL"}
	pkg.SaveToCSV(posts, "bitsearch_scrape.csv", headers)

	elapsedTime := time.Since(startTime)
	fmt.Printf("Waktu eksekusi: %s\n", elapsedTime)
}
