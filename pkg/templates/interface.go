package templates

import (
	"bytes"
	"text/template"
	htpl "html/template"

)

type Tpl interface {
	//Execute applies a parsed template to the specified data object,
	Execute(tplName string, text string, data interface{}) (string, error)
}

//TextTpl is the representation of the text parsed template
type TextTpl struct {
	template *template.Template
}

func NewTextTpl() Tpl {
	tt := &TextTpl{}

	tt.template = &template.Template{}

	return tt
}


func (tt TextTpl) Execute(tplName string, text string, data interface{}) (string, error) {
	tmpl, err := tt.template.New(tplName).Funcs(EsimFuncMap()).
		Parse(text)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

//HtmlTpl is the representation of the html parsed template
type HtmlTpl struct {
	template *htpl.Template
}

func NewHtmlTpl() Tpl {
	tt := &HtmlTpl{}

	tt.template = &htpl.Template{}

	return tt
}

func (ht HtmlTpl) Execute(tplName string, text string, data interface{}) (string, error) {
	tmpl, err := ht.template.New(tplName).Funcs(EsimFuncMap()).
		Parse(text)
	if err != nil {
		return "", err
	}

	var buf bytes.Buffer
	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}