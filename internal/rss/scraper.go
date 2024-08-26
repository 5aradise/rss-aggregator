package rss

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/5aradise/rss-aggregator/internal/db"
	"github.com/google/uuid"
)

func StartScraping(
	ctx context.Context,
	dbQ *db.Queries,
	conurentRequests int32,
	timeToRequest time.Duration,
	timeBetweenRequests time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration\n", conurentRequests, timeBetweenRequests)
	t := time.NewTicker(timeBetweenRequests)
	wg := &sync.WaitGroup{}
	for ; ; <-t.C {
		feeds, err := dbQ.GetNextFeedsToFetch(ctx, conurentRequests)
		if err != nil {
			log.Printf("Error getting next feeds to fetch: %v\n", err)
			continue
		}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapUrl(ctx, wg, dbQ, feed, timeToRequest)
		}
		wg.Wait()
	}
}

func scrapUrl(ctx context.Context, wg *sync.WaitGroup, dbQ *db.Queries, dbFeed db.Feed, ttr time.Duration) {
	defer wg.Done()

	_, err := dbQ.MarkFeedAsFetched(ctx, dbFeed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v\n", err)
		return
	}

	feed, err := urlToFeed(dbFeed.Url, ttr)
	if err != nil {
		log.Printf("Error getting feed from url: %v\n", err)
		return
	}

	newPostsCount := 0

	for _, item := range feed.Channel.Item {
		pubAt, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("Couldn't parse date %s with err %v\n", item.PubDate, err)
			continue
		}

		desc := sql.NullString{}
		if item.Description != "" {
			desc.String = item.Description
			desc.Valid = true
		}

		now := time.Now().UTC()
		_, err = dbQ.CreatPost(ctx, db.CreatPostParams{
			ID:          uuid.New(),
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       item.Title,
			Description: desc,
			PublishedAt: pubAt,
			Url:         item.Link,
			FeedID:      dbFeed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}

			log.Printf("Failed to create post: %v\n", err)
			continue
		}

		newPostsCount++
	}

	log.Printf("Feed %s collected. New %d posts found\n", dbFeed.Name, newPostsCount)
}
