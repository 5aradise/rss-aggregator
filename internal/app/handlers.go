package app

import (
	"github.com/go-chi/chi/v5"
)

func setHandlers(s *Server) {
	v1R := chi.NewRouter()
	v1R.Get("/healthz", handlerReadiness)

	s.mux.Mount("/v1", v1R)
}
