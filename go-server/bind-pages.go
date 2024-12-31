package main

import (
	"net/http"
	"os"
	"text/template"
)

func BindPages(mux *http.ServeMux) {
	mux.HandleFunc("GET /", fa_index_page)
	mux.HandleFunc("GET /aine", aine_page)
	mux.HandleFunc("GET /peyvand", peyvand_page)
}

func fa_index_page(w http.ResponseWriter, r *http.Request) {

	templ := PageTemplate{}
	templ.Name = "index"
	templ.Language = "fa"

	path := "../pages/index/"

	// template head info
	templ.Title = "تربتی | Torbatti"
	templ.Author = "آریا بختیاری | Arya Bakhtiari"
	templ.Description = "وبسایت شخصی آریا بختیاری"

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
	t, err := template.New(templ.Name).Parse(templ.String())
	if err != nil {
		w.Write([]byte("creating template error:" + err.Error()))
	}

	// Executing Template
	err = t.ExecuteTemplate(w, templ.Name, nil)
	if err != nil {
		w.Write([]byte("template execution error:" + err.Error()))
	}
}

func aine_page(w http.ResponseWriter, r *http.Request) {

	templ := PageTemplate{}
	templ.Name = "mirror"
	templ.Language = "fa"

	path := "../pages/aine/"

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
	t, err := template.New(templ.Name).Parse(templ.String())
	if err != nil {
		w.Write([]byte("creating template error:" + err.Error()))
	}

	// Executing Template
	err = t.ExecuteTemplate(w, templ.Name, nil)
	if err != nil {
		w.Write([]byte("template execution error:" + err.Error()))
	}
}
func peyvand_page(w http.ResponseWriter, r *http.Request) {

	templ := PageTemplate{}
	templ.Name = "mirror"
	templ.Language = "fa"

	path := "../pages/peyvand/"

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
	t, err := template.New(templ.Name).Parse(templ.String())
	if err != nil {
		w.Write([]byte("creating template error:" + err.Error()))
	}

	// Executing Template
	err = t.ExecuteTemplate(w, templ.Name, nil)
	if err != nil {
		w.Write([]byte("template execution error:" + err.Error()))
	}
}
