package htmltemplates

import (
	"fmt"
	"io"
)

type ViewData struct {
	Strings []string
}

func Render(w io.Writer, data ViewData) error {
	var result string
	result += "<ul>"
	for _, s := range data.Strings {
		result += fmt.Sprintf("<li>%s</li>", s)
	}
	result += "</ul>"
	_, err := fmt.Fprint(w, result)
	return err
}
