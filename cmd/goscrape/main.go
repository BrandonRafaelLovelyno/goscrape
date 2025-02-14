package main

import (
	"log"

	"github.com/BrandonRafaelLovelyno/goscrape/internal/cli"
	"github.com/BrandonRafaelLovelyno/goscrape/internal/scraping"
	"github.com/BrandonRafaelLovelyno/goscrape/pkg/json"
)

func main() {
	cmdInput, err := cli.ParseCommandInput()
	if err != nil {
		log.Fatalf("Failed to parse user input: %v", err)
	}

	header := scraping.NewScraperHeader("Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Mobile Safari/537.36", nil)
	scraper := scraping.NewScraper(cmdInput.Url, header, cmdInput)

	node, err := scraper.Scrape()
	if err != nil {
		log.Fatalf("Failed to scrape: %v", err)
	}

	jsonData, err := json.ConvertToJson(node)
	if err != nil {
		log.Fatalf("Failed to convert node to json: %v", err)
	}

	err = json.WriteToJson(jsonData, cmdInput.OutDir)
	if err != nil {
		log.Fatalf("Failed to write to json: %v", err)
	}

	log.Printf("Scraping completed, output written to %s", cmdInput.OutDir)
}
