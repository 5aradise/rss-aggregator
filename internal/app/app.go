package app

import (
	"log"
	"net"
	"net/http"

	"github.com/5aradise/rss-aggregator/config"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func Run(cfg config.Config) error {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1R := chi.NewRouter()
	v1R.Get("/healthz", handlerReadiness)

	r.Mount("/v1", v1R)

	srv := &http.Server{
		Addr:    net.JoinHostPort("", cfg.Server.Port),
		Handler: r,
	}

	log.Printf("Starting HTTP server on port %s", cfg.Server.Port)
	return srv.ListenAndServe()
}
