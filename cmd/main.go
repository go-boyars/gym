package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/go-boyars/gym/internal/config"
	"github.com/go-boyars/gym/internal/exercise-app"
	"github.com/go-boyars/gym/internal/exercise-app/pg"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	cfg := config.GetConfig(os.Getenv("CONFIG_PATH"))
	config.Secret = cfg.TokenSign

	pool, err := pgxpool.Connect(context.Background(), cfg.DBconn)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	pg, err := pg.NewPgRepository(pool)
	if err != nil {
		log.Fatal(err)
	}

	app, err := exercise.New(pg)
	if err != nil {
		log.Fatal(err)
	}

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: app.Router(),
	}

	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal(err)
	}
}
