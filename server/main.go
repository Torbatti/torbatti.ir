package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Create Server Mux
	r := http.NewServeMux()

	// Check healthyness
	r.HandleFunc("GET /health", health_check)

	// bind
	BindPublic(r)
	BindMirror(r)

	// BindIndex(r)
	// BindContent(r)

	// create server struct
	server := http.Server{
		Addr:    ":8585",
		Handler: r,
	}
	log.Println("Starting server on port", "localhost"+server.Addr)

	//listen and serve
	server.ListenAndServe()
}

func health_check(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
