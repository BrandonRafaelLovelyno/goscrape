package main

import (
	"log"

	"github.com/BrandonRafaelLovelyno/goscrape/internal/cli"
	"github.com/BrandonRafaelLovelyno/goscrape/internal/scraping"
	"github.com/BrandonRafaelLovelyno/goscrape/pkg/json"
)

func main() {
	arg, err := cli.GetArguments()
	if err != nil {
		log.Fatalf("failed to get command arguments: %v", err)
	}

	scraper := scraping.NewScraper(arg.Url)

	jsonData, err := scraper.Scrape()
	if err != nil {
		log.Fatalf("failed to scrape: %v", err)
	}

	err = json.WriteToJson(jsonData, arg.Output)
	if err != nil {
		log.Fatalf("failed to write to json: %v", err)
	}

	log.Printf("Result exported to: %v", arg.Output)
}
