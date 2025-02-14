package scraping

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func parseWholeHtml(doc *goquery.Document, root *Node) {
	parseNodeChildren(doc.Selection, root)
}

func (s *Scraper) parseWithSelectors(doc *goquery.Document, root *Node) {
	for _, selector := range *s.targetSelectors {
		doc.Find(selector).Each(func(i int, el *goquery.Selection) {
			node := makeNode(el)
			root.Children = append(root.Children, node)
			parseNodeChildren(el, node)
		})
	}
}

func parseNodeChildren(element *goquery.Selection, parent *Node) {
	element.Children().Each(func(i int, el *goquery.Selection) {
		childNode := makeNode(el)
		parent.Children = append(parent.Children, childNode)
		parseNodeChildren(el, childNode)
	})
}

func makeNode(node *goquery.Selection) *Node {
	return &Node{
		ID:       node.AttrOr("id", ""),
		Tag:      node.Get(0).Data,
		Text:     node.Text(),
		Rel:      parseAtrrToSlice(node, "rel"),
		Href:     node.AttrOr("href", ""),
		Alt:      node.AttrOr("alt", ""),
		Title:    node.AttrOr("title", ""),
		Class:    parseAtrrToSlice(node, "class"),
		Name:     node.AttrOr("name", ""),
		Value:    node.AttrOr("value", ""),
		Children: make([]*Node, 0),
	}
}

func parseAtrrToSlice(node *goquery.Selection, attr string) []string {
	attrs := strings.Split(node.AttrOr(attr, ""), " ")
	if len(attrs) == 1 && attrs[0] == "" {
		return nil
	}

	return attrs
}
