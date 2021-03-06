package components

import (
	"html/template"
	"github.com/chenhg5/go-admin/template/types"
)

type ColAttribute struct {
	Name    string
	Content template.HTML
	Size    string
}


func (compo *ColAttribute) SetContent(value template.HTML) types.ColAttribute {
	compo.Content = value
	return compo
}

func (compo *ColAttribute) SetSize(value map[string]string) types.ColAttribute {
	compo.Size = ""
	for key, size := range value {
		compo.Size += "col-" + key + "-" + size + " "
	}
	return compo
}

func (compo *ColAttribute) GetContent() template.HTML {
	return ComposeHtml(*compo, "col")
}