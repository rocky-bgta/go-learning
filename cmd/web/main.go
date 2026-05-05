package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rocky_bgta/go-learning/pkg/config"
	"github.com/rocky_bgta/go-learning/pkg/handlers"
	"github.com/rocky_bgta/go-learning/pkg/render"
)

const portNumber = ":8080"

// main is the main function
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannot create template cache")
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	fmt.Println(fmt.Sprintf("Starting application on port %s", portNumber))

	srv := &http.Server{
		Addr:    portNumber,
		Handler: router(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
