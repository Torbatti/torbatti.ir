package main

import "net/http"

func BindPublic(mux *http.ServeMux) {
	// Serve Static Files
	public_static := http.FileServer(http.Dir("../public"))
	mux.Handle("GET /public/", http.StripPrefix("/public/", public_static))
}
