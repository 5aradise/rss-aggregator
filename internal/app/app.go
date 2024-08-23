package app

import (
	"context"
	"database/sql"
	"log"
	"net"
	"net/http"

	"github.com/5aradise/rss-aggregator/config"
	"github.com/5aradise/rss-aggregator/internal/db"
	"github.com/5aradise/rss-aggregator/internal/rss"
	"github.com/go-chi/chi/v5"

	_ "github.com/lib/pq"
)

type App struct {
	db *db.Queries
}

func Run(cfg config.Config) error {
	conn, err := sql.Open("postgres", cfg.DB.URL)
	if err != nil {
		return err
	}

	app := App{
		db: db.New(conn),
	}

	go rss.StartScraping(
		context.Background(),
		app.db,
		cfg.RSS.ConurentRequests,
		cfg.RSS.TimeToRequest,
		cfg.RSS.TimeBetweenRequests,
	)

	r := chi.NewRouter()

	app.setMiddlewares(r)
	app.setHandlers(r)

	srv := &http.Server{
		Addr:    net.JoinHostPort("", cfg.Server.Port),
		Handler: r,
	}

	log.Printf("Starting HTTP server on port %s", cfg.Server.Port)
	return srv.ListenAndServe()
}
