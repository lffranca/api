package export

import (
	"bytes"
	"html/template"

	"github.com/lffranca/api/export/html"
)

// NewHTMLInput NewHTMLInput
type NewHTMLInput struct {
	Template []byte
	Input    interface{}
	FuncMap  template.FuncMap
}

// NewHTML NewHTML
func NewHTML(input *NewHTMLInput) (Export, error) {
	t, errT := template.New("").Option("missingkey=zero").Funcs(input.FuncMap).Parse(string(input.Template))
	if errT != nil {
		return nil, errT
	}

	tpl := new(bytes.Buffer)
	if err := t.Execute(tpl, input.Input); err != nil {
		return nil, err
	}

	return &html.HTML{
		Buffer: tpl,
	}, nil
}
