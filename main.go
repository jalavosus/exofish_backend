package main

import (
	"exofish-backend/handlers"
	"fmt"
	"log"
	"net/http"
)

/**
 * The root/apex of backend.exo.fish has no use, but it's still fun to have the route work.
 */
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to rootland")
}

func main() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/shuttles", handlers.BookShuttleHandler)
	http.HandleFunc("/shuttleTimes", handlers.GetTimesHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
