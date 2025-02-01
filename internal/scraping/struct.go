package scraping

import "github.com/go-rod/rod"

type Scraper struct {
	url       string
	header    *Header
	page      *rod.Page
	waitNodes []string
}

type Header struct {
	userAgent string
	cookie    []Cookie
}

type Cookie struct {
	Name  string
	Value string
}
