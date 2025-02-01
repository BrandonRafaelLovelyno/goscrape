package main

import (
	"log"

	"github.com/BrandonRafaelLovelyno/goscrape/internal/scraping"
	"github.com/BrandonRafaelLovelyno/goscrape/pkg/json"
)

func main() {
	scraper := scraping.NewScraper("https://proxyway.com/guides/best-websites-to-practice-your-web-scraping-skills")

	jsonData, err := scraper.Scrape()
	if err != nil {
		log.Fatalf("failed to scrape: %v", err)
	}

	err = json.WriteToJson(jsonData, "output.json")
	if err != nil {
		log.Fatalf("failed to write to json: %v", err)
	}
}
