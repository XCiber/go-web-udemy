package main

import (
	"fmt"
	"github.com/XCiber/go-web-udemy/pkg/config"
	"github.com/XCiber/go-web-udemy/pkg/handlers"
	"github.com/XCiber/go-web-udemy/pkg/render"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const port = 8080

// main is the app entry point
func main() {

	var app config.AppConfig
	app.UseCache = true

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatalf("Can't get templates cache: %v", err)
	}
	app.TemplateCache = tc

	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplate(&app)

	log.Printf("Starting app server on port %d", port)

	go func() {
		srv := &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: routes(&app),
		}
		err = srv.ListenAndServe()
		if err != nil {
			log.Fatalf("Could not start server on port %d: %v", port, err)
		}
	}()

	// Wait until some signal is captured.
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, syscall.SIGTERM, syscall.SIGINT)
	<-sigC
}
