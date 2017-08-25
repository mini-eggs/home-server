package app

import (
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// Start - start app
func Start() {
	router := mux.NewRouter()
	server := &http.Server{
		Handler:      router,
		Addr:         ":5000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	registerRoutes(router)

	log.Fatal(server.ListenAndServe())
}
