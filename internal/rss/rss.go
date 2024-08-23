package rss

import (
	"encoding/xml"
	"io"
	"net/http"
	"time"
)

type Feed struct {
	Channel struct {
		Title       string `xml:"title"`
		Link        string `xml:"link"`
		Description string `xml:"description"`
		Language    string `xml:"language"`
		Item        []Item `xml:"item"`
	} `xml:"channel"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func urlToFeed(url string, ttr time.Duration) (Feed, error) {
	client := http.Client{
		Timeout: ttr,
	}

	resp, err := client.Get(url)
	if err != nil {
		return Feed{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return Feed{}, err
	}

	feed := Feed{}
	err = xml.Unmarshal(dat, &feed)
	if err != nil {
		return Feed{}, err
	}

	return feed, nil
}
