package htmltemplates

import (
	"embed"
	"html/template"
	"io"
)

type ViewData struct {
	Strings []string
	Author  string
}

type ViewDataRenderer struct {
	templ *template.Template
}

var (
	//go:embed templates/*
	viewTemplate embed.FS
)

func NewRenderer() (*ViewDataRenderer, error) {
	templ, err := template.ParseFS(viewTemplate, "templates/*.gohtml")
	if err != nil {
		return nil, err
	}
	return &ViewDataRenderer{templ: templ}, nil
}

func (r *ViewDataRenderer) Render(w io.Writer, v ViewData) error {
	if err := r.templ.Execute(w, v); err != nil {
		return err
	}
	return nil
}
