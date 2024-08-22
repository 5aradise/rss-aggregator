package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/5aradise/rss-aggregator/internal/db"
	"github.com/5aradise/rss-aggregator/internal/entities"
	"github.com/5aradise/rss-aggregator/pkg/req"
	"github.com/5aradise/rss-aggregator/pkg/resp"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func (app *App) createFeedFollow(w http.ResponseWriter, r *http.Request, user entities.User) {
	type reqParams struct {
		FeedID uuid.UUID `json:"feed_id"`
	}

	params := reqParams{}
	err := req.DecodeJSON(r, &params)
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	now := time.Now().UTC()
	feedFollow, err := app.db.CreateFeedFollow(r.Context(), db.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn`t create feed follow: %v", err))
		return
	}

	resp.WithJSON(w, http.StatusCreated, entities.ConvertDbFeedFollow(feedFollow))
}

func (app *App) listFeedFollows(w http.ResponseWriter, r *http.Request) {
	feedFollows, err := app.db.ListFeedsFollows(r.Context())
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp.WithJSON(w, http.StatusOK, entities.ConvertDbFeedFollows(feedFollows))
}

func (app *App) getFeedFollows(w http.ResponseWriter, r *http.Request, user entities.User) {
	feedFollows, err := app.db.ListFeedFollowsByUserID(r.Context(), user.ID)
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, err.Error())
	}

	resp.WithJSON(w, http.StatusCreated, entities.ConvertDbFeedFollows(feedFollows))
}

func (app *App) deleteFeedFollow(w http.ResponseWriter, r *http.Request, user entities.User) {
	ffIDStr := chi.URLParam(r, "feedFollowID")
	feedFollow, err := uuid.Parse(ffIDStr)
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing UUID: %v", err))
		return
	}

	err = app.db.DeleteFeedFollow(r.Context(), db.DeleteFeedFollowParams{
		ID:     feedFollow,
		UserID: user.ID,
	})
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp.WithJSON(w, http.StatusOK, struct{}{})
}
