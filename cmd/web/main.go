package main

import (
	"fmt"
	"github.com/XCiber/go-web-udemy/pkg/handlers"
	"log"
	"net/http"
)

const port = 8080

// main is the app entry point
func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	log.Printf("Starting app server on port %d", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		log.Fatalf("Could not start server on port %d: %v", port, err)
	}

}
