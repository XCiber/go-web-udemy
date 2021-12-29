package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/XCiber/go-web-udemy/pkg/config"
	"github.com/XCiber/go-web-udemy/pkg/handlers"
	"github.com/XCiber/go-web-udemy/pkg/render"
	"github.com/alexedwards/scs/v2"
)

const port = 8080

var app config.AppConfig
var session *scs.SessionManager

// main is the app entry point
func main() {

	// change it to true in production
	app.InProduction = true

	app.UseCache = true

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session

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
