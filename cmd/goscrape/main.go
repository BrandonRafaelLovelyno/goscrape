package main

import (
	"log"

	"github.com/BrandonRafaelLovelyno/goscrape/internal/cli"
	"github.com/BrandonRafaelLovelyno/goscrape/internal/scraping"
	"github.com/BrandonRafaelLovelyno/goscrape/pkg/json"
)

func main() {

	log.Println("Getting Command Arguments")
	arg, err := cli.GetAllArguments()
	if err != nil {
		log.Fatalf("failed to get command arguments: %v", err)
	}

	header := scraping.NewScraperHeader("Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Mobile Safari/537.36", nil)
	scraper := scraping.NewScraper(arg.Url, header, arg)

	node, err := scraper.Scrape()
	if err != nil {
		log.Fatalf("failed to scrape: %v", err)
	}

	jsonData, err := json.ConvertToJson(node)
	if err != nil {
		log.Fatalf("failed to convert to json: %v", err)
	}

	err = json.WriteToJson(jsonData, arg.OutDir)
	if err != nil {
		log.Fatalf("failed to write to json: %v", err)
	}

	log.Printf("Scraping completed, output written to %s", arg.OutDir)
}
