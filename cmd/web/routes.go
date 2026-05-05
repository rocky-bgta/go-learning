package main

import (
	"net/http"

	"github.com/bmizerany/pat"
	"github.com/rocky_bgta/go-learning/pkg/config"
	"github.com/rocky_bgta/go-learning/pkg/handlers"
)

func router(app *config.AppConfig) http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))

	return mux

}
