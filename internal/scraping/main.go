package scraping

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/proto"
)

func NewScraper(url string, header *Header) *Scraper {
	return &Scraper{
		url:       url,
		header:    header,
		waitNodes: []string{"body"},
	}
}

func NewScraperHeader(userAgent string, cookies []Cookie) *Header {
	return &Header{
		userAgent: userAgent,
		cookie:    cookies,
	}
}

func (s *Scraper) Scrape() (*[]byte, error) {
	html, err := s.GetHtml()
	if err != nil {
		return nil, fmt.Errorf("failed to get page HTML", err.Error())
	}

	return nil, nil
}

func (s *Scraper) GetHtml() (*string, error) {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(s.url)
	defer page.MustClose()

	err := s.addHeader()
	if err != nil {
		return nil, fmt.Errorf("failed to add header: ", err.Error())
	}

	s.waitData()

	html, err := page.HTML()
	if err != nil {
		return nil, err
	}

	return &html, nil
}

func (s *Scraper) addHeader() error {
	cookies := s.getCookies()

	s.page.MustSetCookies(cookies...)

	err := s.setUserAgent()
	if err != nil {
		return err
	}

	return nil
}

func (s *Scraper) waitData() {
	for _, node := range s.waitNodes {
		s.page.MustElement(node).MustWaitVisible()
	}
}
