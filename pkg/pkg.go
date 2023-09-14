package pkg

import (
	"detik-scraper-gocolly/models"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

func SaveToCSV(posts interface{}, fileName string, headers []string) {
	currentTime := time.Now()
	fileNameWithTime := currentTime.Format("2006_01_02_15_04_") + fileName

	dirName := "results"

	if err := ensureDirectoryExists(dirName); err != nil {
		log.Fatal(err)
	}

	filePath := filepath.Join(dirName, fileNameWithTime)

	file, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Menulis header ke dalam file CSV
	writer.Write(headers)

	// Menulis data post ke dalam file CSV
	switch v := posts.(type) {
	case []models.Btdig:
		for _, post := range v {
			row := []string{post.Date, post.Title, post.Files, post.Size, post.Link, post.MagnetURL}
			writer.Write(row)
		}
	case []models.BitSearch:
		for _, post := range v {
			row := []string{post.Date, post.Title, post.Size, post.Seeder, post.Leecher, post.Downloader, post.Link, post.TorrentLink, post.MagnetURL}
			writer.Write(row)
		}
	}

	fmt.Println("Data post telah disimpan dalam", fileNameWithTime)
}

func ensureDirectoryExists(directory string) error {
	// Mengecek apakah direktori sudah ada
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		// Jika belum ada, maka buat direktori tersebut
		err := os.MkdirAll(directory, os.ModePerm)
		if err != nil {
			return err
		}
	}
	return nil
}
