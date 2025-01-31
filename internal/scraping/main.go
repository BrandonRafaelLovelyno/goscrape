package scraping

import (
	"github.com/gocolly/colly/v2"
)

func NewScraper(url string) *Scraper {
	return &Scraper{
		url: url,
		root: &Node{
			Tag: "html",
		},
	}
}

func (s *Scraper) Start() {
	c := colly.NewCollector()
	s.addCallback(c)
	c.Visit(s.url)
}

func (s *Scraper) addCallback(c *colly.Collector) {
	c.OnHTML("body *", func(e *colly.HTMLElement) {
		node := parseDOM(e)

		parentKey, err := getParentKey(e)
		if err != nil {
			return
		}

		parentNode := s.nodes[parentKey]
		s.appendParentChildren(parentNode, node)

		s.addToNodes(e, node)
	})
}
