package scraping

import (
	"fmt"
	"log"
	"strings"

	"github.com/gocolly/colly/v2"
)

func parseDOM(e *colly.HTMLElement) *Node {
	return &Node{
		Tag:      e.Name,
		Class:    parseAtrrToSlice(e, "class"),
		ID:       e.Attr("id"),
		Text:     e.Text,
		Rel:      parseAtrrToSlice(e, "rel"),
		Src:      e.Attr("src"),
		Href:     e.Attr("href"),
		Alt:      e.Attr("alt"),
		Title:    e.Attr("title"),
		Name:     e.Attr("name"),
		Value:    e.Attr("value"),
		Children: make([]*Node, 0),
	}
}

func (s *Scraper) appendNodeAsChildren(parentNode *Node, node *Node) {
	parentNode.Text = ""
	parentNode.Children = append(parentNode.Children, node)
}

func (s *Scraper) addToNodes(e *colly.HTMLElement, node *Node) {
	nodeKey := fmt.Sprintf("%p", e.DOM.Get(0))
	s.nodes[nodeKey] = node
}

func getParentKey(e *colly.HTMLElement) (string, error) {
	parent := e.DOM.Parent().Get(0)
	if parent == nil {
		return "", fmt.Errorf("no parent found")
	}

	parentKey := fmt.Sprintf("%p", parent)
	return parentKey, nil
}

func parseAtrrToSlice(e *colly.HTMLElement, attr string) []string {
	return strings.Split(e.Attr(attr), " ")
}

func addErrorCallback(c *colly.Collector) {
	c.OnError(func(r *colly.Response, err error) {
		log.Printf("error: %v", err)

		err = r.Request.Retry()
		if err != nil {
			log.Printf("retry failed: %v", err)
		}
	})
}

func addResponseCallback(c *colly.Collector) {
	c.OnResponse(func(r *colly.Response) {
		log.Printf("Response code %d received for URL: %s", r.StatusCode, r.Request.URL)
	})
}
