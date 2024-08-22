package app

import (
	"github.com/5aradise/rss-aggregator/internal/auth"
	"github.com/go-chi/chi/v5"
)

func (app *App) setHandlers(mux *chi.Mux) {
	v1R := chi.NewRouter()
	v1R.Get("/healthz", handlerReadiness)

	v1R.Post("/users", app.createUser)
	v1R.Get("/users/list", app.listUsers)
	v1R.Get("/users", auth.Middleware(app.db, app.getUser))

	v1R.Post("/feeds", auth.Middleware(app.db, app.createFeed))
	v1R.Get("/feeds", app.listFeeds)

	v1R.Post("/feed_follows", auth.Middleware(app.db, app.createFeedFollow))
	v1R.Get("/feed_follows/list", app.listFeedFollows)
	v1R.Get("/feed_follows", auth.Middleware(app.db, app.getFeedFollows))
	v1R.Delete("/feed_follows/{feedFollowID}", auth.Middleware(app.db, app.deleteFeedFollow))

	mux.Mount("/v1", v1R)
}
