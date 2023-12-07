package pkg

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"
	"torrent-scraper/models"
)

type Files struct {
	Title string
	Size  int64
}

func GetFileResults() []Files {
	folderPath := "results" // Ganti dengan path ke folder "result" Anda

	files, err := ioutil.ReadDir(folderPath)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		// return
	}

	// Slice untuk menyimpan informasi file
	var fileInfos []Files

	fmt.Println("Files in the 'result' folder:")
	for _, file := range files {
		if file.IsDir() {
			fmt.Println("Directory:", file.Name())
		} else {
			// Menambahkan informasi file ke dalam slice fileInfos
			fileInfo := Files{
				Title: file.Name(),
				Size:  file.Size(),
			}
			fileInfos = append(fileInfos, fileInfo)
			fmt.Println("File:", file.Name(), "| Size:", file.Size(), "bytes")
		}
	}

	// Menampilkan total jumlah file
	fmt.Println("Total number of files:", len(fileInfos))

	// Menampilkan informasi file dari slice fileInfos
	fmt.Println("\nFile Information:")
	for _, fileInfo := range fileInfos {
		fmt.Printf("File: %s | Size: %d bytes\n", fileInfo.Title, fileInfo.Size)
	}

	return fileInfos
}

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
