package app

import (
	"github.com/go-chi/chi/v5"
)

func (app *App) setHandlers(mux *chi.Mux) {
	v1R := chi.NewRouter()
	v1R.Get("/healthz", handlerReadiness)

	v1R.Post("/users", app.createUser)
	v1R.Get("/users", app.listUsers)

	mux.Mount("/v1", v1R)
}
