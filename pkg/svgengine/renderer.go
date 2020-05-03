package svgengine

import (
	"bytes"
	"html/template"
)

func Render(req SvgRequest) (*bytes.Buffer, error) {
	var returnValue bytes.Buffer
	funcMap := template.FuncMap{
		"inc": func(i int) int {
			return i + 1
		},
		"dec": func(i int) int {
			return i - 1
		},
	}
	temp, err := template.New("tracked-svg.svg").
		Funcs(funcMap).
		Parse(TEMPLATE)
	if err != nil {
		return &returnValue, err
	}
	err = temp.ExecuteTemplate(&returnValue, "tracked-svg.svg", req)
	if err != nil {
		return &returnValue, err
	}
	return &returnValue, nil
}
