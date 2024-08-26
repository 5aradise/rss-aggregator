package entities

import (
	"time"

	"github.com/5aradise/rss-aggregator/internal/db"
	"github.com/google/uuid"
)

type Post struct {
	ID          uuid.UUID `jspn:"id"`
	CreatedAt   time.Time `jspn:"created_at"`
	UpdatedAt   time.Time `jspn:"updated_at"`
	Title       string    `jspn:"title"`
	Description string    `jspn:"description"`
	PublishedAt time.Time `jspn:"published_at"`
	Url         string    `jspn:"url"`
	FeedID      uuid.UUID `jspn:"feed_id"`
}

func ConvertDbPost(dbPost db.Post) Post {
	return Post{
		dbPost.ID,
		dbPost.CreatedAt,
		dbPost.UpdatedAt,
		dbPost.Title,
		dbPost.Description.String,
		dbPost.PublishedAt,
		dbPost.Url,
		dbPost.FeedID,
	}
}

func ConvertDbPosts(dbPosts []db.Post) []Post {
	Posts := make([]Post, 0, len(dbPosts))
	for _, dbPost := range dbPosts {
		Posts = append(Posts, ConvertDbPost(dbPost))
	}
	return Posts
}
