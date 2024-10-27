package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func fa_content_shell_hx(w http.ResponseWriter, r *http.Request) {
	root_path := "../fa-content/"
	hub_pv := r.PathValue("hub")
	if hub_pv == "" {
		w.Write([]byte("ERROR: invalid hub_pv: empty hub_pv"))
	}

	shell, err := ContentReadHtml(root_path, hub_pv, "hubshell")
	if err != nil {
		w.Write([]byte("reading file error:" + err.Error()))
	}

	fmt.Println("shell: " + hub_pv)

	w.Header().Set("Cache-Control", "private, max-age=60") // keep cache for 60seconds

	t, err := template.New("contentShow").Parse(`{{define "T"}}` + fa_base_shell + shell + `{{end}}`)
	err = t.ExecuteTemplate(w, "T", nil)
}

func fa_content_hub_hx(w http.ResponseWriter, r *http.Request) {
	root_path := "../fa-content/"

	hub_pv := r.PathValue("hub")
	if hub_pv == "" {
		w.Write([]byte("ERROR: invalid hub_pv: empty hub_pv"))
	}

	hub_html, err := ContentReadHtml(root_path, hub_pv, "hub")
	if err != nil {
		w.Write([]byte("reading file error:" + err.Error()))
	}
	hub_turtle_html, err := ContentReadHtml(root_path, hub_pv, "hubturtle")
	if err != nil {
		w.Write([]byte("reading file error:" + err.Error()))
	}

	var hubshell_force = `<div hx-get="/hx/fa/content/shell/` + hub_pv + `" hx-trigger="load" hx-target="#shell" hx-swap="innerHtml"></div>`

	fmt.Println("hx: " + hub_pv)

	w.Header().Set("Cache-Control", "private, max-age=60") // keep cache for 60seconds
	w.Header().Set("HX-Push-Url", "/fa/content/"+hub_pv)   // push the url

	t, err := template.New("contentShow").Parse(`{{define "T"}}` + hubshell_force + hub_turtle_html + hub_html + `{{end}}`)
	err = t.ExecuteTemplate(w, "T", nil)
}

func fa_content_sub_hx(w http.ResponseWriter, r *http.Request) {

	root_path := "../fa-content/"

	hub_pv := r.PathValue("hub")
	if hub_pv == "" {
		w.Write([]byte("ERROR: invalid hub_pv: empty hub_pv"))
	}
	sub_pv := r.PathValue("sub")
	if sub_pv == "" {
		w.Write([]byte("ERROR: invalid sub_pv: empty sub_pv"))
	}

	sub_html, err := ContentReadHtml(root_path, hub_pv, "sub-"+sub_pv)
	if err != nil {
		w.Write([]byte("reading file error:" + err.Error()))
	}
	hub_turtle_html, err := ContentReadHtml(root_path, hub_pv, "hubturtle")
	if err != nil {
		w.Write([]byte("reading file error:" + err.Error()))
	}

	fmt.Println("hx: " + hub_pv + " - " + sub_pv)

	w.Header().Set("Cache-Control", "private, max-age=60")          // keep cache for 60seconds
	w.Header().Set("HX-Push-Url", "/fa/content/"+hub_pv+"/"+sub_pv) // push the url

	t, err := template.New("contentShow").Parse(`{{define "T"}}` + hub_turtle_html + sub_html + `{{end}}`)
	err = t.ExecuteTemplate(w, "T", nil)
}

func en_content_shell_hx(w http.ResponseWriter, r *http.Request) {
	root_path := "../en-content/"
	hub_pv := r.PathValue("hub")
	if hub_pv == "" {
		w.Write([]byte("ERROR: invalid hub_pv: empty hub_pv"))
	}

	shell, err := ContentReadHtml(root_path, hub_pv, "hubshell")
	if err != nil {
		w.Write([]byte("reading file error:" + err.Error()))
	}

	fmt.Println("shell: " + hub_pv)

	w.Header().Set("Cache-Control", "private, max-age=60") // keep cache for 60seconds

	t, err := template.New("contentShow").Parse(`{{define "T"}}` + en_base_shell + shell + `{{end}}`)
	err = t.ExecuteTemplate(w, "T", nil)
}

func en_content_hub_hx(w http.ResponseWriter, r *http.Request) {
	root_path := "../en-content/"

	hub_pv := r.PathValue("hub")
	if hub_pv == "" {
		w.Write([]byte("ERROR: invalid hub_pv: empty hub_pv"))
	}

	hub_html, err := ContentReadHtml(root_path, hub_pv, "hub")
	if err != nil {
		w.Write([]byte("reading file error:" + err.Error()))
	}
	hub_turtle_html, err := ContentReadHtml(root_path, hub_pv, "hubturtle")
	if err != nil {
		w.Write([]byte("reading file error:" + err.Error()))
	}

	var hubshell_force = `<div hx-get="/hx/en/content/shell/` + hub_pv + `" hx-trigger="load" hx-target="#shell" hx-swap="innerHtml"></div>`

	fmt.Println("hx: " + hub_pv)

	w.Header().Set("Cache-Control", "private, max-age=60") // keep cache for 60seconds
	w.Header().Set("HX-Push-Url", "/en/content/"+hub_pv)   // push the url

	t, err := template.New("contentShow").Parse(`{{define "T"}}` + hubshell_force + hub_turtle_html + hub_html + `{{end}}`)
	err = t.ExecuteTemplate(w, "T", nil)
}

func en_content_sub_hx(w http.ResponseWriter, r *http.Request) {

	root_path := "../en-content/"

	hub_pv := r.PathValue("hub")
	if hub_pv == "" {
		w.Write([]byte("ERROR: invalid hub_pv: empty hub_pv"))
	}
	sub_pv := r.PathValue("sub")
	if sub_pv == "" {
		w.Write([]byte("ERROR: invalid sub_pv: empty sub_pv"))
	}

	sub_html, err := ContentReadHtml(root_path, hub_pv, "sub-"+sub_pv)
	if err != nil {
		w.Write([]byte("reading file error:" + err.Error()))
	}
	hub_turtle_html, err := ContentReadHtml(root_path, hub_pv, "hubturtle")
	if err != nil {
		w.Write([]byte("reading file error:" + err.Error()))
	}

	fmt.Println("hx: " + hub_pv + " - " + sub_pv)

	w.Header().Set("Cache-Control", "private, max-age=60")          // keep cache for 60seconds
	w.Header().Set("HX-Push-Url", "/en/content/"+hub_pv+"/"+sub_pv) // push the url

	t, err := template.New("contentShow").Parse(`{{define "T"}}` + hub_turtle_html + sub_html + `{{end}}`)
	err = t.ExecuteTemplate(w, "T", nil)
}
