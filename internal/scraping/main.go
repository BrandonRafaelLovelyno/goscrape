package scraping

import (
	"fmt"

	"github.com/go-rod/rod"
)

func NewScraper(url string, header *Header) *Scraper {
	return &Scraper{
		url:       url,
		header:    header,
		waitNodes: []string{"body"},
	}
}

func (s *Scraper) GetHtml() (*string, error) {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(s.url)
	defer page.MustClose()

	s.addHeader()
	s.waitData()

	html, err := page.HTML()
	if err != nil {
		return nil, fmt.Errorf("failed to get page html: %v", err.Error())
	}

	return &html, nil
}

func (s *Scraper) addHeader() {
	cookies := s.getCookies()

	s.page.MustSetCookies(cookies...)
}

func NewScraperHeader(userAgent string, cookies []Cookie) *Header {
	return &Header{
		userAgent: userAgent,
		cookie:    cookies,
	}
}
