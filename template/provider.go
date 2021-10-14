package template

import (
	"bytes"
	"html/template"
)

type HTMLTemplateInterface interface {
	Compile(interface{}, string) (bytes.Buffer, error)
}

// HTMLTemplate contains the HTML templates used for PDF content
type HTMLTemplate struct {
	Templates *template.Template
}

// App won't start if the system can't parse the templates.
func NewHTMLTemplate() *HTMLTemplate {
	return &HTMLTemplate{
		Templates: template.Must(template.ParseGlob("./templates/*")),
	}
}

func (ht *HTMLTemplate) Compile(data interface{}, name string) (bytes.Buffer, error) {
	var b bytes.Buffer

	err := ht.Templates.ExecuteTemplate(&b, name, data)
	if err != nil {
		return *new(bytes.Buffer), err
	}

	return b, nil
}
