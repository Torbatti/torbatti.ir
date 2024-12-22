package main

import "net/http"

func BindMirror(mux *http.ServeMux) {
	// Serve Static Files
	mirror_static := http.FileServer(http.Dir("../mirror"))
	mux.Handle("GET /mirror/", http.StripPrefix("/mirror/", mirror_static))
}
