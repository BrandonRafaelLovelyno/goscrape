package scraping

import "github.com/gocolly/colly/v2"

type Scraper struct {
	url string
	c   *colly.Collector
}

type Node struct {
	Tag      string   `json:"tag"`
	Class    []string `json:"class,omitempty"`
	Rel      []string `json:"rel,omitempty"`
	Src      string   `json:"src,omitempty"`
	Href     string   `json:"href,omitempty"`
	Alt      string   `json:"alt,omitempty"`
	Title    string   `json:"title,omitempty"`
	Name     string   `json:"name,omitempty"`
	Value    string   `json:"value,omitempty"`
	ID       string   `json:"id,omitempty"`
	Text     string   `json:"text,omitempty"`
	Children []*Node  `json:"children,omitempty"`
}
