package htmltemplates_test

import (
	"HTML_Templates/htmltemplates"
	"bytes"
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

	t.Run("it converts []string into HTML", func(t *testing.T) {
		buf := bytes.Buffer{}
		err := htmltemplates.Render(&buf, test)
		if err != nil {
			t.Fatal(err)
		}
		approvals.VerifyString(t, buf.String())
	})

}
