package main

import "net/http"

func BindIndex(mux *http.ServeMux) {
	mux.HandleFunc("/", fa_index_page)
	mux.HandleFunc("/en", en_index_page)
}
