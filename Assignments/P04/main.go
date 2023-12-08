package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"
)

// Sequential version of the image downloader.
func downloadImagesSequential(urls []string) {
	for _, url := range urls {
		filename := generateFilename(url)
		err := downloadImage(url, filename)
		if err != nil {
			fmt.Printf("Error downloading %s: %v\n", url, err)
		}
	}
}

// Concurrent version of the image downloader.
func downloadImagesConcurrent(urls []string) {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			filename := generateFilename(u)
			err := downloadImage(u, filename)
			if err != nil {
				fmt.Printf("Error downloading %s: %v\n", u, err)
			}
		}(url)
	}

	wg.Wait()
}

// Helper function to generate a unique filename based on the URL.
func generateFilename(url string) string {
	ext := filepath.Ext(url)
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
}

// Helper function to download and save a single image.
func downloadImage(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	urls := []string{
		"https://unsplash.com/photos/hvdnff_bieQ/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://unsplash.com/photos/HQaZKCDaax0/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		"https://images.unsplash.com/photo-1698778573682-346d219402b5?ixlib=rb-4.0.3&q=85&fm=jpg&crop=entropy&cs=srgb&w=640",
		"https://unsplash.com/photos/Bs2jGUWu4f8/download?ixid=M3wxMjA3fDB8MXx0b3BpY3x8NnNNVmpUTFNrZVF8fHx8fDJ8fDE2OTg5MDc1MDh8&w=640",
		// Add more image URLs
	}

	// Sequential download
	start := time.Now()
	downloadImagesSequential(urls)
	fmt.Printf("Sequential download took: %v\n", time.Since(start))

	// Concurrent download
	start = time.Now()
	downloadImagesConcurrent(urls)
	fmt.Printf("Concurrent download took: %v\n", time.Since(start))
}
