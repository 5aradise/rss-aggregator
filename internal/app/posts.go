package app

import (
	"net/http"
	"strconv"

	"github.com/5aradise/rss-aggregator/internal/db"
	"github.com/5aradise/rss-aggregator/internal/entities"
	"github.com/5aradise/rss-aggregator/pkg/resp"
)

const (
	defaultPostLimit = 10
	minPostLimit     = 1
	maxPostLimit     = 100
)

func (app *App) listPostsForUser(w http.ResponseWriter, r *http.Request, user entities.User) {
	limit := defaultPostLimit

	queryParams := r.URL.Query()
	limitStr := queryParams.Get("limit")
	if limitStr != "" {
		limitFromQuery, err := strconv.Atoi(limitStr)
		if err == nil {
			if limitFromQuery < 1 {
				limit = minPostLimit
			} else if limitFromQuery > 100 {
				limit = maxPostLimit
			} else {
				limit = limitFromQuery
			}
		}
	}

	dbPosts, err := app.db.ListPostsForUser(r.Context(), db.ListPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp.WithJSON(w, http.StatusOK, entities.ConvertDbPosts(dbPosts))
}
