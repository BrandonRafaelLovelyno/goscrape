package scraping

import (
	"fmt"

	"github.com/go-rod/rod/lib/proto"
)

func (s *Scraper) addHeader() error {
	cookies := s.getCookies()

	if len(cookies) > 0 {
		s.page.MustSetCookies(cookies...)
	}

	err := s.setUserAgent()
	if err != nil {
		return fmt.Errorf("failed to set user agent: %v", err.Error())
	}

	return nil
}

func (s *Scraper) getCookies() []*proto.NetworkCookieParam {
	cookies := make([]*proto.NetworkCookieParam, 0)

	if s.header.cookie == nil {
		return cookies
	}

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
