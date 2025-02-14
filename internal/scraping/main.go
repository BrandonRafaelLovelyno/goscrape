package scraping

import (
	"fmt"

	"github.com/BrandonRafaelLovelyno/goscrape/internal/cli"
)

func NewScraper(url string, header *Header, cmdInput *cli.CommandInput) *Scraper {
	return &Scraper{
		url:             url,
		header:          header,
		waitedSelectors: cmdInput.WaitedSelectors,
		targetSelectors: cmdInput.TargetSelectors,
	}
}

func NewScraperHeader(userAgent string, cookies []Cookie) *Header {
	return &Header{
		userAgent: userAgent,
		cookie:    cookies,
	}
}

func (s *Scraper) Scrape() (*Node, error) {
	html, err := s.getHtml()
	if err != nil {
		return nil, fmt.Errorf("failed to get page HTML: %v", err.Error())
	}

	doc, err := s.readHtmlDocument(html)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTML document: %v", err.Error())
	}

	root := s.parseHtmlDocument(doc)

	return root, nil
}
