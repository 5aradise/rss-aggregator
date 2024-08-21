package entities

import (
	"time"

	"github.com/5aradise/rss-aggregator/internal/db"
	"github.com/google/uuid"
)

type Feed struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Url       string    `json:"url"`
	UserID    uuid.UUID `json:"user_id"`
}

func ConvertDbFeed(dbFeed db.Feed) Feed {
	return Feed{
		dbFeed.ID,
		dbFeed.CreatedAt,
		dbFeed.UpdatedAt,
		dbFeed.Name,
		dbFeed.Url,
		dbFeed.UserID,
	}
}

func ConvertDbFeeds(dbFeeds []db.Feed) []Feed {
	feeds := make([]Feed, 0, len(dbFeeds))
	for _, dbFeed := range dbFeeds {
		feeds = append(feeds, ConvertDbFeed(dbFeed))
	}
	return feeds
}
