package main

func CreateFaPageTemplate() PageTemplateType {
	layout := LayoutType{
		HtmlStart: fa_html_start,
		HeadStart: fa_head_start,
		HeadEnd:   fa_head_end,
		BodyStart: fa_body_start,
		BodyEnd:   fa_body_end,
		HtmlEnd:   fa_html_end,

		ShellStart:  fa_shell_start,
		Shellend:    fa_shell_end,
		TurtleStart: fa_turtle_start,
		Turtleend:   fa_turtle_end,
	}

	page := PageType{
		HeadExtend:  "",
		BodyExtend:  "",
		ShellExtend: "",
	}
	head_info := HeadInfoType{
		Title:       "",
		Author:      "",
		Description: "",
	}

	return PageTemplateType{
		LayoutName: "",
		Layout:     layout,
		Page:       page,
		HeadInfo:   head_info,
	}
}
func CreateEnPageTemplate() PageTemplateType {
	layout := LayoutType{
		HtmlStart: en_html_start,
		HeadStart: en_head_start,
		HeadEnd:   en_head_end,
		BodyStart: en_body_start,
		BodyEnd:   en_body_end,
		HtmlEnd:   en_html_end,

		ShellStart:  en_shell_start,
		Shellend:    en_shell_end,
		TurtleStart: en_turtle_start,
		Turtleend:   en_turtle_end,
	}

	page := PageType{
		HeadExtend:  "",
		BodyExtend:  "",
		ShellExtend: "",
	}
	head_info := HeadInfoType{
		Title:       "",
		Author:      "",
		Description: "",
	}

	return PageTemplateType{
		LayoutName: "",
		Layout:     layout,
		Page:       page,
		HeadInfo:   head_info,
	}
}
