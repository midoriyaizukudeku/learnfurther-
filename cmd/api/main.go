package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log/slog"
	"main/internal/data"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn         string
		maxOpenConn int
		maxIdleConn int
		MaxIdleTime time.Duration
	}
}

type application struct {
	config config
	logger *slog.Logger
	models data.Models
}

func main() {

	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "PORT NETWORM")
	flag.StringVar(&cfg.env, "env", "development", "ENVIzroment(deve|staging|production)")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("GREENLIGHT_DB_DSN_NOSSL"), "POSTQRESl DSN")
	flag.IntVar(&cfg.db.maxOpenConn, "db-maxopenconn", 25, "maximum open connections")
	flag.IntVar(&cfg.db.maxIdleConn, "db-maxIdlecon", 25, "maximum idle connections")
	flag.DurationVar(&cfg.db.MaxIdleTime, "db-maxIdleTime", 15*time.Minute, "Maximum Idle timeout ")

	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	db, err := OpenDb(cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	logger.Info("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/healthcheck", app.healthcheckHandler)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      app.Routes(),
		IdleTimeout:  1 * time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	logger.Info("starting server", "addr", srv.Addr, "env", cfg.env)
	err = srv.ListenAndServe()
	logger.Error(err.Error())
	os.Exit(1)
}

func OpenDb(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(cfg.db.maxOpenConn)
	db.SetMaxIdleConns(cfg.db.maxIdleConn)
	db.SetConnMaxIdleTime(cfg.db.MaxIdleTime)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
