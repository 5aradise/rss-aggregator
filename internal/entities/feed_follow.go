package entities

import (
	"time"

	"github.com/5aradise/rss-aggregator/internal/db"
	"github.com/google/uuid"
)

type FeedFollow struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	UserID    uuid.UUID `json:"user_id"`
	FeedID    uuid.UUID `json:"feed_id"`
}

func ConvertDbFeedFollow(dbFeed db.FeedFollow) FeedFollow {
	return FeedFollow{
		dbFeed.ID,
		dbFeed.CreatedAt,
		dbFeed.UpdatedAt,
		dbFeed.UserID,
		dbFeed.FeedID,
	}
}

func ConvertDbFeedFollows(dbFeedFollows []db.FeedFollow) []FeedFollow {
	feedFollows := make([]FeedFollow, 0, len(dbFeedFollows))
	for _, dbFeedFollow := range dbFeedFollows {
		feedFollows = append(feedFollows, ConvertDbFeedFollow(dbFeedFollow))
	}
	return feedFollows
}
