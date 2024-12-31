package main

import (
	"net/http"
	"os"
	"text/template"
)

func BindContent(mux *http.ServeMux) {
	mux.HandleFunc("GET /nevise/{hub}", nevise_hub_page)
	mux.HandleFunc("GET /nevise/{hub}/{sub}", nevise_sub_page)
}

func nevise_hub_page(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "private, max-age="+CacheTime) // keep cache for 60seconds

	// url: /fa/nevise/{hub}
	hub_value := r.PathValue("hub")
	if hub_value == "" {
		w.Write([]byte("ERROR: invalid hub_value: empty hub_pv"))
	}

	templ := PageTemplate{}
	templ.Name = "hub"
	templ.Language = "fa"

	path := "../content/" + hub_value + "/"

	// template head
	templ.ExtendHead = ""
	head_html, err := os.ReadFile(path + "head" + ".html")
	if err != nil {
		w.Write([]byte("ERROR-head_html : " + err.Error()))
	}
	templ.ExtendHead = templ.ExtendHead + string(head_html)

	// template body
	templ.ExtendBody = ""
	body_html, err := os.ReadFile(path + "body" + ".html")
	if err != nil {
		w.Write([]byte("ERROR-index_html : " + err.Error()))
	}
	templ.ExtendBody = templ.ExtendBody + string(body_html)

	// Creating Template
	t, err := template.New("neviseShow").Parse(templ.String())
	if err != nil {
		w.Write([]byte("creating template error:" + err.Error()))
	}

	// Executing Template
	err = t.ExecuteTemplate(w, templ.Name, nil)
	if err != nil {
		w.Write([]byte("template execution error:" + err.Error()))
	}
}

func nevise_sub_page(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Cache-Control", "private, max-age="+CacheTime) // keep cache for 60seconds

	// url: /fa/nevise/{hub}/{sub}
	hub_value := r.PathValue("hub")
	if hub_value == "" {
		w.Write([]byte("ERROR: invalid hub_value: empty hub_value"))
	}
	sub_value := r.PathValue("sub")
	if sub_value == "" {
		w.Write([]byte("ERROR: invalid sub_value: empty sub_value"))
	}

	templ := PageTemplate{}
	templ.Name = "sub"
	templ.Language = "fa"

	path := "../content/" + hub_value + "/" + sub_value + "/"

	// template head
	templ.ExtendHead = ""
	head_html, err := os.ReadFile(path + "head" + ".html")
	if err != nil {
		w.Write([]byte("ERROR-head_html : " + err.Error()))
	}
	templ.ExtendHead = templ.ExtendHead + string(head_html)

	// template body
	templ.ExtendBody = ""
	body_html, err := os.ReadFile(path + "body" + ".html")
	if err != nil {
		w.Write([]byte("ERROR-index_html : " + err.Error()))
	}
	templ.ExtendBody = templ.ExtendBody + string(body_html)

	// Creating Template
	t, err := template.New("neviseShow").Parse(templ.String())
	if err != nil {
		w.Write([]byte("creating template error:" + err.Error()))
	}

	// Executing Template
	err = t.ExecuteTemplate(w, templ.Name, nil)
	if err != nil {
		w.Write([]byte("template execution error:" + err.Error()))
	}
}
