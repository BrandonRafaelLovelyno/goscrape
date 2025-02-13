package scraping

import (
	"fmt"
	"strings"

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

func (s *Scraper) Scrape() (*Node, error) {
	html, err := s.GetHtml()
	if err != nil {
		return nil, fmt.Errorf("failed to get page HTML", err.Error())
	}

	doc, err := s.readHtmlDocument(html)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTML document", err.Error())
	}

	root := parseHtmlDocument(doc)

	return root, nil
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

func (s *Scraper) readHtmlDocument(html *string) (*goquery.Document, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(*html))
	if err != nil {
		return nil, err
	}

	return doc, nil
}

func parseHtmlDocument(doc *goquery.Document) *Node {
	root := &Node{Tag: "root", Children: make([]*Node, 0)}
	parseNode(doc.Selection, root)

	return root
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
