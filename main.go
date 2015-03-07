package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/johnweldon/fmtf/formatter"
)

func main() {
	ff := formatter.NewFormatter()
	if err := format(os.Stdin, os.Stdout, ff); err != nil {
		fmt.Fprintf(os.Stdout, "error formatting: %q\n", err)
		os.Exit(-1)
	}
}

func format(reader io.Reader, writer io.Writer, ff formatter.Formatter) error {
	var (
		buf, dst []byte
		err      error
	)
	if buf, err = ioutil.ReadAll(reader); err != nil {
		fmt.Fprintf(writer, "%s", buf)
		return fmt.Errorf("Error reading source: %v", err)
	}
	if dst, err = ff.Filter(buf); err != nil {
		fmt.Fprintf(writer, "%s", buf)
		return fmt.Errorf("Error indenting: %v", err)
	}
	fmt.Fprintf(writer, "%s", dst)
	return nil
}
