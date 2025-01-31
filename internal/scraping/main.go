package scraping

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func NewScraper(url string) *Scraper {
	return &Scraper{
		url: url,
		root: &Node{
			Tag: "html",
		},
		nodes: make(map[string]*Node),
	}
}

func (s *Scraper) Start() error {
	c := colly.NewCollector()

	addVerboseCallback(c)
	s.addHtmlCallback(c)

	c.Visit(s.url)

	jsonData, _ := s.getJson()
	log.Println(jsonData)
	return nil
}

func addVerboseCallback(c *colly.Collector) {
	addErrorCallback(c)
	addResponseCallback(c)
}

func (s *Scraper) addHtmlCallback(c *colly.Collector) {
	c.OnHTML("body *", func(e *colly.HTMLElement) {
		if strings.Contains(e.Name, "script") {
			return
		}

		node := parseDOM(e)

		parentKey, err := getParentKey(e)
		if err != nil {
			log.Println("Error on getting parent key:", err)
			return
		}

		parentNode, exists := s.nodes[parentKey]
		if !exists {
			parentNode = s.root
			s.nodes[parentKey] = s.root
		}

		s.appendNodeAsChildren(parentNode, node)

		s.addToNodes(e, node)
	})
}

func (s *Scraper) getJson() (string, error) {
	jsonData, err := json.MarshalIndent(s.root, "", "  ")
	if err != nil {
		return "", fmt.Errorf("failed to indent marshaling Json: %v", err)
	}

	return string(jsonData), nil
}
