package rssgen

import (
	"bytes"
	"fmt"
	"github.com/gorilla/feeds"
	"mvdan.cc/xurls"
	"strings"
	"time"
)

func reformattedHtml(link string) (html string, err error) {
	link = strings.TrimSpace(link)
	if len(link) < 5 {
		return "", fmt.Errorf("Wrong link.")
	}
	lowerLink := strings.ToLower(link)

	schemeIndex := strings.Index(lowerLink, "http://")
	if schemeIndex == -1 {
		schemeIndex = strings.Index(lowerLink, "https://")
	}

	if schemeIndex == -1 {
		var buffer bytes.Buffer
		buffer.WriteString("http://")
		buffer.WriteString(strings.TrimLeft(link, ":/"))
		link = buffer.String()
		lowerLink = strings.ToLower(link)
	}

	_, html, err = getHttp(link)
	if err != nil {
		return "", err
	}

	html = reformatHtml(link, html)
	return html, nil
}

func generateRSS(generator Generator) (rss string, err error) {
	html, err := reformattedHtml(generator.Url)
	if err != nil {
		return "", err
	}

	gen := GenRegex{Rule: generator.Rule}
	feedItems, size, err := gen.Parse(html)
	if err != nil {
		return "", err
	}

	now := time.Now()
	feed := &feeds.Feed{
		Title:       generator.Title,
		Link:        &feeds.Link{Href: generator.Url},
		Description: generator.Description,
		Created:     now,
	}

	feed.Items = make([]*feeds.Item, 0, size)
	for _, item := range feedItems {
		item["link"] = xurls.Strict().FindString(item["link"])
		if len(item["link"]) > 0 && len(item["title"]) > 0 {
			feed.Items = append(feed.Items, &feeds.Item{
				Id:          item["link"],
				Title:       item["title"],
				Link:        &feeds.Link{Href: item["link"]},
				Author:      &feeds.Author{Name: item["name"], Email: item["email"]},
				Description: item["description"],
				Created:     now,
			})
		}
	}

	rss, err = feed.ToAtom()
	if err != nil {
		return "", err
	}

	return rss, nil
}
