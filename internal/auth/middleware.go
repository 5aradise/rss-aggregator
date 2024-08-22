package auth

import (
	"net/http"

	"github.com/5aradise/rss-aggregator/internal/db"
	"github.com/5aradise/rss-aggregator/internal/entities"
	"github.com/5aradise/rss-aggregator/pkg/resp"
)

type authedHandler func(http.ResponseWriter, *http.Request, entities.User)

func Middleware(db *db.Queries, h authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := GetApiKey(r.Header)
		if err != nil {
			resp.WithError(w, http.StatusUnauthorized, "Bad request: "+err.Error())
			return
		}

		user, err := db.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			resp.WithError(w, http.StatusUnauthorized, "Bad request: "+err.Error())
			return
		}

		h(w, r, entities.ConvertDbUser(user))
	}
}
