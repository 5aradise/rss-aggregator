package app

import (
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/5aradise/rss-aggregator/config"
	"github.com/5aradise/rss-aggregator/internal/db"
	"github.com/go-chi/chi/v5"

	_ "github.com/lib/pq"
)

type App struct {
	db *db.Queries
}

func Run(cfg config.Config) error {
	conn, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		return err
	}

	app := App{
		db: db.New(conn),
	}

	r := chi.NewRouter()

	app.setMiddlewares(r)
	app.setHandlers(r)

	srv := &http.Server{
		Addr:    net.JoinHostPort("", cfg.Port),
		Handler: r,
	}

	log.Printf("Starting HTTP server on port %s", cfg.Port)
	return srv.ListenAndServe()
}
