package main

import "net/http"

func BindContent(mux *http.ServeMux) {
	//fa
	mux.HandleFunc("GET /hx/fa/content/{hub}", fa_content_hub_hx)
	mux.HandleFunc("GET /hx/fa/content/{hub}/{sub}", fa_content_sub_hx)

	mux.HandleFunc("GET /hx/fa/content/shell/{hub}", fa_content_shell_hx)

	mux.HandleFunc("GET /fa/content/{hub}", fa_content_hub_page)
	mux.HandleFunc("GET /fa/content/{hub}/{sub}", fa_content_sub_page)

	//en
	mux.HandleFunc("GET /hx/en/content/{hub}", en_content_hub_hx)
	mux.HandleFunc("GET /hx/en/content/{hub}/{sub}", en_content_sub_hx)

	mux.HandleFunc("GET /hx/en/content/shell/{hub}", en_content_shell_hx)

	mux.HandleFunc("GET /en/content/{hub}", en_content_hub_page)
	mux.HandleFunc("GET /en/content/{hub}/{sub}", en_content_sub_page)
}
