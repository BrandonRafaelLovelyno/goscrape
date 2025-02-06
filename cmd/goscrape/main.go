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

<<<<<<< Updated upstream
	header := scraping.NewScraperHeader("Mozilla/5.0 (Linux; Android 10; K) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/127.0.0.0 Mobile Safari/537.36", nil)
	scraper := scraping.NewScraper(arg.Url, header)
=======
	scraper := scraping.NewScraper(arg.Url)

	jsonData, err := scraper.Scrape()
	if err != nil {
		log.Fatalf("failed to scrape: %v", err)
	}

	err = json.WriteToJson(jsonData, arg.OutDir)
	if err != nil {
		log.Fatalf("failed to write to json: %v", err)
	}

	log.Printf("Result exported to: %v", arg.OutDir)
>>>>>>> Stashed changes
}
