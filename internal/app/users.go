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

func (app *App) createUser(w http.ResponseWriter, r *http.Request) {
	type reqParams struct {
		Name string `json:"name"`
	}

	params := reqParams{}
	err := req.DecodeJSON(r, &params)
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, fmt.Sprintf("Error parsing JSON: %v", err))
		return
	}

	now := time.Now().UTC()
	user, err := app.db.CreateUser(r.Context(), db.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      params.Name,
	})
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, fmt.Sprintf("Couldn`t create user: %v", err))
		return
	}

	resp.WithJSON(w, http.StatusCreated, entities.ConvertDbUser(user))
}

func (app *App) listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := app.db.ListUsers(r.Context())
	if err != nil {
		resp.WithError(w, http.StatusBadRequest, err.Error())
		return
	}

	resp.WithJSON(w, http.StatusOK, entities.ConvertDbUsers(users))
}

func (app *App) getUser(w http.ResponseWriter, r *http.Request, user entities.User) {
	resp.WithJSON(w, http.StatusOK, user)
}
