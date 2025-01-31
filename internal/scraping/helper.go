package scraping

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly/v2"
)

func parseDOM(e *colly.HTMLElement) *Node {
	return &Node{
		Tag:   e.Name,
		Class: parseAtrrToSlice(e, "class"),
		ID:    e.Attr("id"),
		Text:  e.Text,
		Rel:   parseAtrrToSlice(e, "rel"),
		Src:   e.Attr("src"),
		Href:  e.Attr("href"),
		Alt:   e.Attr("alt"),
		Title: e.Attr("title"),
		Name:  e.Attr("name"),
		Value: e.Attr("value"),
	}
}

func (s *Scraper) appendParentChildren(parentNode *Node, node *Node) {
	if parentNode.Children == nil {
		parentNode.Children = make(map[string][]*Node)
	}

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
