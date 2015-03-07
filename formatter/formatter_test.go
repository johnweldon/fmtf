package formatter_test

import (
	"testing"

	"github.com/johnweldon/fmtf/formatter"
)

var testCase = map[string]string{
	`{"key1":"val1","key2":"val2"}`: "{\n  \"key1\": \"val1\",\n  \"key2\": \"val2\"\n}",
	"<root><hi>asdf</hi></root>":    "<root>\n  <hi>asdf</hi>\n</root>",
	"asdfasdf":                      "asdfasdf",
}

func TestBasic(t *testing.T) {
	fr := formatter.NewFormatter()
	for given, expect := range testCase {
		out, err := fr.Filter([]byte(given))
		if err != nil {
			t.Errorf("unexpected error %q", err)
		}
		if expect != string(out) {
			t.Errorf("mismatch:\n   given %q\nexpected %q\n     got %q\n", given, expect, out)
		}
	}
}
