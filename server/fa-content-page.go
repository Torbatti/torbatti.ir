package main

import (
	"html/template"
	"net/http"
	"os"
)

func fa_content_hub_page(w http.ResponseWriter, r *http.Request) {
	// url: /fa/content/{hub}
	hub_value := r.PathValue("hub")
	if hub_value == "" {
		w.Write([]byte("ERROR: invalid hub_value: empty hub_pv"))
	}

	templ := PageTemplate{}
	templ.Name = "hub"
	templ.Language = "fa"

	path := "../fa/content/" + hub_value + "/"

	// template head
	head_html, err := os.ReadFile(path + "head" + ".html")
	if err != nil {
		w.Write([]byte("ERROR-head_html : " + err.Error()))
	}
	templ.ExtendHead = string(head_html)

	// template body
	templ.ExtendBody = "" // will not be used when turtle-shell is available

	// template turtle-shell
	shell_html, err := os.ReadFile(path + "shell" + ".html")
	if err != nil {
		w.Write([]byte("ERROR-shell_html : " + err.Error()))
	}
	turtle_html, err := os.ReadFile(path + "turtle" + ".html")
	if err != nil {
		w.Write([]byte("ERROR-turtle_html : " + err.Error()))
	}
	templ.Turtle = string(turtle_html)
	templ.Shell = string(shell_html)

	// Creating Template
	t, err := template.New("contentShow").Parse(templ.String())
	if err != nil {
		w.Write([]byte("creating template error:" + err.Error()))
	}

	// Executing Template
	err = t.ExecuteTemplate(w, templ.Name, nil)
	if err != nil {
		w.Write([]byte("template execution error:" + err.Error()))
	}
}

func fa_content_sub_page(w http.ResponseWriter, r *http.Request) {
	// url: /fa/content/{hub}/{sub}
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

	path := "../fa/content/" + hub_value + "/"

	// template head
	head_html, err := os.ReadFile(path + "subhead-" + sub_value + ".html")
	if err != nil {
		w.Write([]byte("ERROR-head_html : " + err.Error()))
	}
	templ.ExtendHead = string(head_html)

	// template body
	templ.ExtendBody = "" // will not be used when turtle-shell is available

	// template turtle-shell
	shell_html, err := os.ReadFile(path + "shell" + ".html")
	if err != nil {
		w.Write([]byte("ERROR-shell_html : " + err.Error()))
	}
	turtle_html, err := os.ReadFile(path + "subturtle-" + sub_value + ".html")
	if err != nil {
		w.Write([]byte("ERROR-turtle_html : " + err.Error()))
	}
	templ.Turtle = string(turtle_html)
	templ.Shell = string(shell_html)

	// Creating Template
	t, err := template.New("contentShow").Parse(templ.String())
	if err != nil {
		w.Write([]byte("creating template error:" + err.Error()))
	}

	// Executing Template
	err = t.ExecuteTemplate(w, templ.Name, nil)
	if err != nil {
		w.Write([]byte("template execution error:" + err.Error()))
	}
}
