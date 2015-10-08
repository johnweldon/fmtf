package formatter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"unicode"

	"github.com/clbanning/mxj"
)

// Formatter can filter incoming bytes and potentially format
// them with indentation and regular style.
type Formatter interface {
	Filter(in []byte) ([]byte, error)
}

// NewFormatter returns a base Formatter that can handle at least
// xml and json.
func NewFormatter() Formatter {
	return &filterMap{
		'"':  formatJSON,
		'\'': formatJSON,
		'[':  formatJSON,
		'{':  formatJSON,
		'<':  formatXML,
	}
}

type filterFn func(in []byte) ([]byte, error)
type filterMap map[rune]filterFn

var (
	nop filterFn = func(in []byte) ([]byte, error) { return in, nil }
	tag filterFn = func(in []byte) ([]byte, error) { return append(in, []byte("xxx")...), nil }
)

func (f *filterMap) Filter(in []byte) ([]byte, error) {
	filter := f.pickFilter(in)
	return filter(in)
}

func (f *filterMap) pickFilter(in []byte) filterFn {
	r, err := findFirstNonSpaceRune(in)
	if err != nil {
		return nop
	}
	if filter, ok := map[rune]filterFn(*f)[r]; ok {
		return filter
	}
	return nop
}

func findFirstNonSpaceRune(in []byte) (rune, error) {
	for _, r := range string(in) {
		if !unicode.IsSpace(r) {
			return r, nil
		}
	}
	return 0, fmt.Errorf("only whitespace")
}

func formatXML(in []byte) ([]byte, error) {
	m, err := mxj.NewMapXmlReader(bytes.NewBuffer(in))
	if err != nil {
		return []byte{}, fmt.Errorf("Error indenting: %v", err)
	}
	return m.XmlIndent("", "  ")
}

func formatJSON(in []byte) ([]byte, error) {
	dst := new(bytes.Buffer)
	if err := json.Indent(dst, in, "", "  "); err != nil {
		return []byte{}, fmt.Errorf("Error indenting: %v", err)
	}
	return dst.Bytes(), nil
}
