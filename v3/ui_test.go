package swaggerui

import "testing"

func TestBox(t *testing.T) {
	b := Box()

	files := b.List()

	t.Log("files: ", files)

	for _, expect := range []string{
		"index.html",
	} {
		if !b.Has(expect) {
			t.Error("file not found: ", expect)
		}
	}
}
