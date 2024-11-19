package main

import "log"

type PageTemplate struct {
	// {{define "name"}}
	Name     string
	Language string

	// Head Info
	Title       string
	Author      string
	Description string

	// Head
	ExtendHead string

	// Body
	ExtendBody string

	// Turtle Shell
	Shell  string
	Turtle string
}

func (templ PageTemplate) String() string {
	// concat string parts together
	var templ_string string

	// SHOULD PANIC when name is empty
	if templ.Name == "" {
		log.Fatal("Error : LayoutName can not be Empty!")
	}
	templ_string = `{{define "` + templ.Name + `"}}`

	templ_string = templ_string + `
	<!DOCTYPE html>
	<html>
	<head>
		<meta charset='utf-8'>
		<meta http-equiv='X-UA-Compatible' content='IE=edge'>
		<meta name='viewport' content='width=device-width, initial-scale=1'>

		<!-- PreLoads -->
		<link rel="preload" href="/public/fonts/Vazirmatn[wght].woff2" as="font" type="font/woff2" crossorigin />

		<!-- Fonts -->

		<!-- Stylesheets -->
		<link rel="stylesheet" type='text/css' media='screen' href='/public/styles/main.css' />

		<!-- Scripts -->
		<script src='/public/libs/htmx.min.js'></script>
		<script src='/public/libs/htmx.preload.js'></script>
	`

	if templ.Title != "" {
		templ_string = templ_string + `<title>` + templ.Title + `</title>`
	}
	if templ.Author != "" {
		templ_string = templ_string + `<meta name="author" content="` + templ.Author + `">`
	}
	if templ.Description != "" {
		templ_string = templ_string + `<meta name="description" content="` + templ.Description + `">`
	}

	if templ.ExtendHead != "" {
		templ_string = templ_string + templ.ExtendHead
	}
	templ_string = templ_string + `
	</head>
	`

	if templ.Language != "" {
		templ_string = templ_string + `
		<body hx-ext="preload" class="lang-` + templ.Language + `" >
		`
	} else {
		templ_string = templ_string + `
		<body hx-ext="preload">
		`
	}

	if templ.Shell != "" || templ.Turtle != "" {
		templ_string = templ_string + `
		<div id="turtle-shell">
		<div id="shell">
		`
		templ_string = templ_string + templ.Shell
		templ_string = templ_string + `
		</div>
    	<div id="turtle">
		`
		templ_string = templ_string + templ.Turtle
		templ_string = templ_string + `
		</div>
		</div>
		`
	} else {
		templ_string = templ_string + templ.ExtendBody
	}

	templ_string = templ_string + `
	</body>
	</html>
	`

	templ_string = templ_string + `{{end}}`

	return templ_string
}
