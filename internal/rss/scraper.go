package rss

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/5aradise/rss-aggregator/internal/db"
)

func StartScraping(
	ctx context.Context,
	db *db.Queries,
	conurentRequests int32,
	timeToRequest time.Duration,
	timeBetweenRequests time.Duration,
) {
	log.Printf("Scraping on %v goroutines every %s duration\n", conurentRequests, timeBetweenRequests)
	t := time.NewTicker(timeBetweenRequests)
	wg := &sync.WaitGroup{}
	for ; ; <-t.C {
		feeds, err := db.GetNextFeedsToFetch(ctx, conurentRequests)
		if err != nil {
			log.Printf("Error getting next feeds to fetch: %v\n", err)
			continue
		}

		for _, feed := range feeds {
			wg.Add(1)
			go scrapUrl(ctx, wg, db, feed, timeToRequest)
		}
		wg.Wait()
	}
}

func scrapUrl(ctx context.Context, wg *sync.WaitGroup, db *db.Queries, dbFeed db.Feed, ttr time.Duration) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(ctx, dbFeed.ID)
	if err != nil {
		log.Printf("Error marking feed as fetched: %v\n", err)
		return
	}

	feed, err := urlToFeed(dbFeed.Url, ttr)
	if err != nil {
		log.Printf("Error getting feed from url: %v\n", err)
		return
	}

	log.Println(feed)
}
