package main

import (
	"log"

	"github.com/BrandonRafaelLovelyno/goscrape/internal/cli"
	"github.com/BrandonRafaelLovelyno/goscrape/internal/scraping"
)

func main() {

	log.Println("Getting Command Arguments")
	arg, err := cli.GetAllArguments()
	if err != nil {
		log.Fatalf("failed to get command arguments: %v", err)
	}

	header := scraping.NewScraperHeader("Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Mobile Safari/537.36", nil)
	scraping.NewScraper(arg.Url, header)
}
