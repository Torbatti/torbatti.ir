package main

import (
	"html/template"
	"net/http"
	"os"
)

func en_index_page(w http.ResponseWriter, r *http.Request) {

	templ := PageTemplate{}
	templ.Name = "index"
	templ.Language = "en"

	path := "../en/page/index/"

	// template head info
	templ.Title = "تربتی | Torbatti"
	templ.Author = "آریا بختیاری | Arya Bakhtiari"
	templ.Description = "وبسایت شخصی آریا بختیاری"

	// template head
	templ.ExtendHead = ""

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
