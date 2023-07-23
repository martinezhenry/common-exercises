package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/newrelic/go-agent/v3/newrelic"
)

var logger log.Logger

func main() {

	logger = log.Logger{}

	if err := newApplication(); err != nil {
		logger.Println("server failed: %w", err)
	}
}

func newApplication() error {

	NEWRELIC_APPNAME := os.Getenv("NEWRELIC_APPNAME")
	NEWRELIC_TOKEN := os.Getenv("NEWRELIC_TOKEN")

	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName(NEWRELIC_APPNAME),
		newrelic.ConfigLicense(NEWRELIC_TOKEN),
		newrelic.ConfigAppLogForwardingEnabled(true),
	)
	if err != nil {
		return err
	}

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	pingHandler := newPingHandler()
	router.Get(newrelic.WrapHandleFunc(app, "/ping", pingHandler))

	if err := http.ListenAndServe(":8080", router); err != nil {
		return err
	}

	return nil
}

func newPingHandler() func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		for i := 0; i < 1000000; i++ {
			fmt.Printf("loop count: %d\n", i)
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("pong"))
	}
}
