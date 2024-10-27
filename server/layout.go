package main

import (
	"fmt"
)

var PageTemplate PageTemplateType

type PageTemplateType struct {
	LayoutName string
	Layout     LayoutType
	Page       PageType
	HeadInfo   HeadInfoType
}

type LayoutType struct {
	HtmlStart string
	HeadStart string
	HeadEnd   string
	BodyStart string
	BodyEnd   string
	HtmlEnd   string

	ShellStart  string
	Shellend    string
	TurtleStart string
	Turtleend   string
}

type PageType struct {
	HeadExtend  string
	BodyExtend  string
	ShellExtend string
}

type HeadInfoType struct {
	Title       string
	Author      string
	Description string
}

func (temp PageTemplateType) String() (string, error) {
	var temp_string string

	// if temp.Page.BodyExtend == "" {
	// 	fmt.Println("Error : Body is Empty!")
	// 	// return "", nil
	// }
	if temp.LayoutName == "" {
		return "", fmt.Errorf("Error : LayoutName can not be Empty!")
	}

	temp_string = temp_string + `{{define "` + temp.LayoutName + `"}}`
	temp_string = temp_string + temp.Layout.HtmlStart

	temp_string = temp_string + temp.Layout.HeadStart
	if temp.HeadInfo.Title != "" {
		temp_string = temp_string + `<title>` + temp.HeadInfo.Title + `</title>`
	}
	if temp.HeadInfo.Author != "" {
		temp_string = temp_string + `<meta name="author" content="` + temp.HeadInfo.Author + `">`
	}
	if temp.HeadInfo.Description != "" {
		temp_string = temp_string + `<meta name="description" content="` + temp.HeadInfo.Description + `">`
	}
	temp_string = temp_string + temp.Page.HeadExtend
	temp_string = temp_string + temp.Layout.HeadEnd

	temp_string = temp_string + temp.Layout.BodyStart
	temp_string = temp_string + temp.Layout.ShellStart
	temp_string = temp_string + temp.Page.ShellExtend
	temp_string = temp_string + temp.Layout.Shellend
	temp_string = temp_string + temp.Layout.TurtleStart
	temp_string = temp_string + temp.Page.BodyExtend
	temp_string = temp_string + temp.Layout.Turtleend
	temp_string = temp_string + temp.Layout.BodyEnd

	temp_string = temp_string + temp.Layout.HtmlEnd
	temp_string = temp_string + `{{end}}`

	return temp_string, nil
}
