package scraping

import (
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

func (s *Scraper) waitData() {
	for _, node := range s.waitNodes {
		s.page.MustElement(node)
	}
}
