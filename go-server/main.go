package main

import (
	"fmt"
	"log"
	"net/http"
)

var CacheTime string

func main() {
	CacheTime = "60"

	// Create Server Mux
	r := http.NewServeMux()

	// Check healthyness
	r.HandleFunc("GET /health", health_check)

	// bind
	BindPublic(r)
	BindMirror(r)

	BindPages(r)
	BindContent(r)

	// create server struct
	server := http.Server{
		Addr:    ":8585",
		Handler: Logging(r),
	}
	log.Println("Starting server on port", "localhost"+server.Addr)

	//listen and serve
	server.ListenAndServe()
}

func health_check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
