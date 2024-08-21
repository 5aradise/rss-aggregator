package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/5aradise/rss-aggregator/internal/db"
	"github.com/5aradise/rss-aggregator/internal/entities"
	"github.com/5aradise/rss-aggregator/pkg/req"
	"github.com/5aradise/rss-aggregator/pkg/resp"
	"github.com/google/uuid"
)

func (app *App) createFeed(w http.ResponseWriter, r *http.Request, user entities.User) {
	type reqParams struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	params := reqParams{}
	err := req.DecodeJSON(r, &params)
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	now := time.Now().UTC()
	feed, err := app.db.CreateFeed(r.Context(), db.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn`t create user: %v", err))
		return
	}

	resp.WithJSON(w, http.StatusCreated, entities.ConvertDbFeed(feed))
}

func (app *App) listFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := app.db.ListFeeds(r.Context())
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn`t create user: %v", err))
		return
	}

	resp.WithJSON(w, http.StatusOK, entities.ConvertDbFeeds(feeds))
}