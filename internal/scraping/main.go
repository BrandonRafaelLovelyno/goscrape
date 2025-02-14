package scraping

import (
	"fmt"
	"strings"

	"github.com/BrandonRafaelLovelyno/goscrape/internal/cli"
	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod"
)

func NewScraper(url string, header *Header, arg *cli.Argument) *Scraper {
	return &Scraper{
		url:             url,
		header:          header,
		waitedSelectors: arg.WaitedSelectors,
		targetSelectors: arg.TargetSelectors,
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
		return nil, fmt.Errorf("failed to get page HTML: %v", err.Error())
	}

	doc, err := s.readHtmlDocument(html)
	if err != nil {
		return nil, fmt.Errorf("failed to read HTML document: %v", err.Error())
	}

	root := s.parseHtmlDocument(doc)

	return root, nil
}

func (s *Scraper) GetHtml() (*string, error) {
	browser := rod.New().MustConnect()
	defer browser.MustClose()

	page := browser.MustPage(s.url)
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
		return nil, err
	}

	return doc, nil
}

func (s *Scraper) parseHtmlDocument(doc *goquery.Document) *Node {
	root := &Node{Tag: "root", Children: make([]*Node, 0)}

	if s.targetSelectors == nil || *s.targetSelectors == nil {
		parseNodeChildren(doc.Selection, root)
		return root
	}

	for _, selector := range *s.targetSelectors {
		doc.Find(selector).Each(func(i int, el *goquery.Selection) {
			node := makeNode(el)
			root.Children = append(root.Children, node)
			parseNodeChildren(el, node)
		})
	}

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
	if s.waitedSelectors == nil || *s.waitedSelectors == nil {
		return
	}

	for _, node := range *s.waitedSelectors {
		s.page.MustElement(node).MustWaitVisible()
	}
}
