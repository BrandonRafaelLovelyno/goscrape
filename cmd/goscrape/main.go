package main

import (
	"log"

	"github.com/BrandonRafaelLovelyno/goscrape/internal/cli"
	"github.com/BrandonRafaelLovelyno/goscrape/internal/scraping"
	"github.com/BrandonRafaelLovelyno/goscrape/pkg/json"
)

func main() {
	url, err := cli.GetArguments()
	if err != nil {
		log.Fatalf("failed to get command arguments: %v", err)
	}

	scraper := scraping.NewScraper(url)

	jsonData, err := scraper.Scrape()
	if err != nil {
		log.Fatalf("failed to scrape: %v", err)
	}

	err = json.WriteToJson(jsonData, "output.json")
	if err != nil {
		log.Fatalf("failed to write to json: %v", err)
	}
}
