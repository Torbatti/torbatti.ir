package main

import (
	"fmt"
	"log"
	"net/http"
)

var root_path string

func main() {
	root_path = "../"

	r := http.NewServeMux()
	// Check healthyness
	r.HandleFunc("GET /health", health)

	// Serve Static Files
	public_static := http.FileServer(http.Dir("../public"))
	r.Handle("/public/", http.StripPrefix("/public/", public_static))

	// Serve Client files
	// client_static := http.FileServer(http.Dir("./client"))
	// r.Handle("/", client_static)

	BindIndex(r)
	BindContent(r)

	// Hx Routes
	// r.HandleFunc("GET /hx/content/list", hx_content_list)

	server := http.Server{
		Addr:    ":8585",
		Handler: r,
	}
	log.Println("Starting server on port", server.Addr)

	server.ListenAndServe()
}

func health(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}

// func hx_content_list(w http.ResponseWriter, r *http.Request) {
// 	path, err := os.Getwd()
// 	if err != nil {
// 		log.Println(err)
// 	}

// 	fileInfos, err := os.ReadDir(path + "/content")
// 	if err != nil {
// 		fmt.Println("Error in accessing directory:", err)
// 	}

// 	var file_names []string
// 	var content_list string
// 	for _, file := range fileInfos {
// 		// fmt.Printf("%T: %+v\n", file, file)
// 		file_names = append(file_names, file.Name())
// 	}
// 	for _, name := range file_names {
// 		content_list = content_list + `<p hx-get="/hx/content/show/` + name + `" hx-trigger="click" hx-target="#content_show" hx-swap="innerHtml">`
// 		content_list = content_list + name
// 		content_list = content_list + "</p>"

// 		// log.Println(name)
// 	}

// 	fmt.Fprintf(w, content_list)
// }
