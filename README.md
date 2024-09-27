# rss-aggregator

## Description

A RESTful web rss aggregator written in Go programming language.

It uses [Chi](https://github.com/go-chi/chi) as the HTTP router and [PostgreSQL](https://www.postgresql.org/) as the database with [sqlc](https://github.com/sqlc-dev/sqlc) and [goose](https://github.com/pressly/goose). It utilizes goroutines for faster feed aggregation.

## Features
- RESTful api
- Concurrent feed aggregation

## Technologies
- Go
- Chi
- Sqlc
- Goose
- PostgreSQL

## Requirements
- Go 1.23
- Goose
- PostgreSQL

## Running the application
1. Clone the repository
2. Create a copy of the `.env.example` file and rename it to `.env`, change `DB_URL` to yours
3. Run `goose postgres <DB_URL> up`
4. Run `make run` to start the server
