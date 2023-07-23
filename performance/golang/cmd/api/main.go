package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var logger log.Logger

func main() {
	logger = log.Logger{}

	if err := newApplication(); err != nil {
		logger.Println("server failed: %w", err)
	}
}

func newApplication() error {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("pong"))
	})

	if err := http.ListenAndServe(":8080", router); err != nil {
		return err
	}

	return nil
}
