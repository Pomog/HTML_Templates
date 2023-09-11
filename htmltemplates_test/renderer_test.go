package htmltemplates_test

import (
	"HTML_Templates/htmltemplates"
	"bytes"
	"testing"
)

func TestRender(t *testing.T) {
	var (
		test = htmltemplates.ViewData{
			Strings: []string{"hello", "world"},
		}
	)

	t.Run("it converts []string into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := htmltemplates.Render(&buf, test)
		if err != nil {
			t.Fatal(err)
		}
		got := buf.String()
		want := `<ul><li>hello</li><li>world</li></ul>`
		if got != want {
			t.Errorf("got '%s' want '%s'", got, want)
		}
	})
}
