package htmltemplates_test

import (
	"HTML_Templates/htmltemplates"
	"bytes"
	"io"
	"testing"

	approvals "github.com/approvals/go-approval-tests"
)

func TestRender(t *testing.T) {
	var (
		test = htmltemplates.ViewData{
			Strings: []string{"hello", "world"},
			Author:  "PomoBlin",
		}
	)

	viewRenderer, err := htmltemplates.NewRenderer()
	if err != nil {
		t.Fatal(err)
	}

	t.Run("it converts []string into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		if err := viewRenderer.Render(&buf, test); err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

}

func BenchmarkRender(b *testing.B) {
	var (
		test = htmltemplates.ViewData{
			Strings: []string{"hello", "world", "this", "is", "a", "test"},
		}
	)

	viewRenderer, err := htmltemplates.NewRenderer()
	if err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		viewRenderer.Render(io.Discard, test)
	}
}
