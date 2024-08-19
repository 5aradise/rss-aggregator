package app

import (
	"log"
	"net"
	"net/http"

	"github.com/5aradise/rss-aggregator/config"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	cfg config.Config
	mux *chi.Mux
}

func New(cfg config.Config) *Server {
	s := &Server{cfg, chi.NewRouter()}

	setMiddlewares(s)
	setHandlers(s)

	return s
}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:    net.JoinHostPort("", s.cfg.Port),
		Handler: s.mux,
	}

	log.Printf("Starting HTTP server on port %s", s.cfg.Port)
	return srv.ListenAndServe()
}
