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

var (
	//go:embed templates/*
	viewTemplate embed.FS
)

func Render(w io.Writer, view ViewData) error {
	templ, err := template.ParseFS(viewTemplate, "templates/*.gohtml")
	if err != nil {
		return err
	}
	if err := templ.Execute(w, view); err != nil {
		return err
	}
	return nil
}
