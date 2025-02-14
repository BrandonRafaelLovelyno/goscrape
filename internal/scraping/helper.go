package scraping

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/go-rod/rod/lib/proto"
)

func (s *Scraper) getCookies() []*proto.NetworkCookieParam {
	cookies := make([]*proto.NetworkCookieParam, 0)

	for _, c := range s.header.cookie {
		cookie := &proto.NetworkCookieParam{
			Name:  c.Name,
			Value: c.Value,
		}
		cookies = append(cookies, cookie)
	}

	return cookies
}

func (s *Scraper) setUserAgent() error {
	err := proto.NetworkSetUserAgentOverride{
		UserAgent: s.header.userAgent,
	}.Call(s.page)

	if err != nil {
		return err
	}

	return nil
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
