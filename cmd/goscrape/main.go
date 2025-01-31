package main

import (
	"github.com/BrandonRafaelLovelyno/goscrape/internal/scraping"
)

func main() {
	scraper := scraping.NewScraper("https://proxyway.com/guides/best-websites-to-practice-your-web-scraping-skills")
	scraper.Start()
}
