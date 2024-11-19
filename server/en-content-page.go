package main

// import (
// 	"fmt"
// 	"html/template"
// 	"net/http"
// )

// func en_content_hub_page(w http.ResponseWriter, r *http.Request) {

// 	content_path := "../en-content/"

// 	hub_pv := r.PathValue("hub")
// 	if hub_pv == "" {
// 		w.Write([]byte("ERROR: invalid hub_pv: empty hub_pv"))
// 	}

// 	hubshell_html, err := ContentReadHtml(content_path, hub_pv, "hubshell")
// 	if err != nil {
// 		w.Write([]byte("reading file error:" + err.Error()))
// 	}
// 	hubturtle_html, err := ContentReadHtml(content_path, hub_pv, "hubturtle")
// 	if err != nil {
// 		w.Write([]byte("reading file error:" + err.Error()))
// 	}
// 	hubhead_html, err := ContentReadHtml(content_path, hub_pv, "hubhead")
// 	if err != nil {
// 		w.Write([]byte("reading file error:" + err.Error()))
// 	}
// 	hub_html, err := ContentReadHtml(content_path, hub_pv, "hub")
// 	if err != nil {
// 		w.Write([]byte("reading file error:" + err.Error()))
// 	}

// 	pt := CreateEnPageTemplate()

// 	pt.LayoutName = "index"

// 	pt.Page.HeadExtend = "<style>body{direction:rtl;}</style>" + hubhead_html
// 	pt.Page.ShellExtend = en_base_shell + hubshell_html
// 	pt.Page.BodyExtend = hubturtle_html + hub_html

// 	pt_string, err := pt.String()
// 	if err != nil {
// 		w.Write([]byte("stringifing template error:" + err.Error()))
// 	}

// 	t, err := template.New("contentShow").Parse(pt_string)
// 	if err != nil {
// 		w.Write([]byte("creating template error:" + err.Error()))
// 	}

// 	err = t.ExecuteTemplate(w, pt.LayoutName, nil)
// 	if err != nil {
// 		w.Write([]byte("template execution error:" + err.Error()))
// 	}
// }

// func en_content_sub_page(w http.ResponseWriter, r *http.Request) {

// 	content_path := "../en-content/"

// 	hub_pv := r.PathValue("hub")
// 	if hub_pv == "" {
// 		w.Write([]byte("ERROR: invalid hub_pv: empty hub_pv"))
// 	}
// 	sub_pv := r.PathValue("sub")
// 	if sub_pv == "" {
// 		w.Write([]byte("ERROR: invalid sub_pv: empty sub_pv"))
// 	}

// 	hubshell_html, err := ContentReadHtml(content_path, hub_pv, "hubshell")
// 	if err != nil {
// 		w.Write([]byte("reading file error:" + err.Error()))
// 	}
// 	hubturtle_html, err := ContentReadHtml(content_path, hub_pv, "hubturtle")
// 	if err != nil {
// 		w.Write([]byte("reading file error:" + err.Error()))
// 	}
// 	subhead_html, err := ContentReadHtml(content_path, hub_pv, "subhead-"+sub_pv)
// 	if err != nil {
// 		// w.Write([]byte("reading file error:" + err.Error()))
// 		fmt.Println("reading file error:" + err.Error())
// 	}
// 	sub_html, err := ContentReadHtml(content_path, hub_pv, "sub-"+sub_pv)
// 	if err != nil {
// 		w.Write([]byte("reading file error:" + err.Error()))
// 	}

// 	pt := CreateEnPageTemplate()

// 	pt.LayoutName = "index"

// 	pt.Page.HeadExtend = "<style>body{direction:rtl;}</style>" + subhead_html
// 	pt.Page.ShellExtend = en_base_shell + hubshell_html
// 	pt.Page.BodyExtend = hubturtle_html + sub_html

// 	pt_string, err := pt.String()
// 	if err != nil {
// 		w.Write([]byte("stringifing template error:" + err.Error()))
// 	}

// 	t, err := template.New("contentShow").Parse(pt_string)
// 	if err != nil {
// 		w.Write([]byte("creating template error:" + err.Error()))
// 	}

// 	err = t.ExecuteTemplate(w, pt.LayoutName, nil)
// 	if err != nil {
// 		w.Write([]byte("template execution error:" + err.Error()))
// 	}
// }
