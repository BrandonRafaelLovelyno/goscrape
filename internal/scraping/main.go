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
		c:   colly.NewCollector(),
		url: url,
	}
}

func (s *Scraper) Scrape() (*[]byte, error) {
	nodes := make(map[string]*Node)
	root := &Node{Tag: "html"}

	s.addVerboseCallback()
	s.addHtmlCallback(nodes, root)

	err := s.visitWebsite()
	if err != nil {
		return nil, fmt.Errorf("failed to visit website: %v", err)
	}

	jsonData, error := s.getJson(root)
	if error != nil {
		return nil, fmt.Errorf("failed to get json: %v", error)
	}

	return jsonData, nil
}

func (s *Scraper) visitWebsite() error {
	err := s.c.Visit(s.url)

	return err
}

func (s *Scraper) addVerboseCallback() {
	addErrorCallback(s.c)
	addResponseCallback(s.c)
}

func (s *Scraper) addHtmlCallback(nodes map[string]*Node, root *Node) {
	s.c.OnHTML("body *", func(e *colly.HTMLElement) {
		if strings.Contains(e.Name, "script") {
			return
		}

		node := parseDOM(e)

		parentKey, err := getParentKey(e)
		if err != nil {
			log.Println("Error on getting parent key:", err)
			return
		}

		parentNode, exists := nodes[parentKey]
		if !exists {
			parentNode = root
			nodes[parentKey] = root
		}

		s.appendNodeAsChildren(parentNode, node)

		s.addToNodes(e, node, nodes)
	})
}

func (s *Scraper) getJson(root *Node) (*[]byte, error) {
	jsonData, err := json.MarshalIndent(root, "", "  ")
	if err != nil {
		return nil, fmt.Errorf("failed to indent marshaling Json: %v", err)
	}

	return &jsonData, nil
}
