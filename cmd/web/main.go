package main

import (
	"fmt"
	"github.com/XCiber/go-web-udemy/pkg/config"
	"github.com/XCiber/go-web-udemy/pkg/handlers"
	"github.com/XCiber/go-web-udemy/pkg/render"
	"log"
	"net/http"
)

const port = 8080

// main is the app entry point
func main() {

	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("Can't get templates cache: %v", err)
	}
	app.TemplateCache = tc
	app.UseCache = false

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)

	log.Printf("Starting app server on port %d", port)
	err = http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Could not start server on port %d: %v", port, err)
	}

}
