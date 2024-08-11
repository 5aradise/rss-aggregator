package app

import (
	"net/http"

	"github.com/5aradise/rss-aggregator/pkg/resp"
)

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	resp.WithJSON(w, http.StatusOK, struct{}{})
}
