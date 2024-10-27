package main

import (
	"html/template"
	"net/http"
	"os"
)

func fa_index_page(w http.ResponseWriter, r *http.Request) {
	pt := CreateFaPageTemplate()

	pt.LayoutName = "index"

	fapage_path := "../fa-pages/"

	indexshell_html, err := os.ReadFile(fapage_path + "indexshell" + ".html")
	if err != nil {
		w.Write([]byte("ERROR-indexshell_html : " + err.Error()))
	}

	pt.HeadInfo.Title = "تربتی | Torbatti"
	pt.HeadInfo.Author = "آریا بختیاری | Arya Bakhtiari"
	pt.HeadInfo.Description = "وبسایت شخصی آریا بختیاری"

	pt.Page.HeadExtend = "<style>body{direction:rtl;}</style>"
	pt.Page.ShellExtend = fa_base_shell + string(indexshell_html)
	pt.Page.BodyExtend = "<div></div>"

	pt_string, err := pt.String()
	if err != nil {
		w.Write([]byte("stringifing template error:" + err.Error()))
	}

	t, err := template.New("contentShow").Parse(pt_string)
	if err != nil {
		w.Write([]byte("creating template error:" + err.Error()))
	}

	err = t.ExecuteTemplate(w, pt.LayoutName, nil)
	if err != nil {
		w.Write([]byte("template execution error:" + err.Error()))
	}

}
func en_index_page(w http.ResponseWriter, r *http.Request) {
	pt := CreateEnPageTemplate()

	pt.LayoutName = "index"

	fapage_path := "../en-pages/"

	indexshell_html, err := os.ReadFile(fapage_path + "indexshell" + ".html")
	if err != nil {
		w.Write([]byte("ERROR-indexshell_html : " + err.Error()))
	}

	pt.HeadInfo.Title = "تربتی | Torbatti"
	pt.HeadInfo.Author = "آریا بختیاری | Arya Bakhtiari"
	pt.HeadInfo.Description = "وبسایت شخصی آریا بختیاری"

	pt.Page.HeadExtend = "<style>body{direction:ltr;}</style>"
	pt.Page.ShellExtend = en_base_shell + string(indexshell_html)
	pt.Page.BodyExtend = "<div></div>"

	pt_string, err := pt.String()
	if err != nil {
		w.Write([]byte("stringifing template error:" + err.Error()))
	}

	t, err := template.New("contentShow").Parse(pt_string)
	if err != nil {
		w.Write([]byte("creating template error:" + err.Error()))
	}

	err = t.ExecuteTemplate(w, pt.LayoutName, nil)
	if err != nil {
		w.Write([]byte("template execution error:" + err.Error()))
	}

}
