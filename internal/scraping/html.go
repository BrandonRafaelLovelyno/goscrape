package scraping

import (
	"fmt"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
)

func (s *Scraper) getHtml() (*string, error) {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(s.url)
	s.page = page
	defer page.MustClose()

	err := s.addHeader()
	if err != nil {
		return nil, fmt.Errorf("failed to add header: %v", err.Error())
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
		return nil, fmt.Errorf("failed to read HTML document: %v", err.Error())
	}

	return doc, nil
}

func (s *Scraper) parseHtmlDocument(doc *goquery.Document) *Node {
	root := &Node{Tag: "root", Children: make([]*Node, 0)}

	if s.targetSelectors == nil {
		parseWholeHtml(doc, root)
	} else {
		s.parseWithSelectors(doc, root)
	}

	return root
}

func (s *Scraper) waitData() {
	if s.waitedSelectors == nil {
		return
	}

	for _, node := range *s.waitedSelectors {
		s.page.MustElement(node).MustWaitVisible()
	}
}
