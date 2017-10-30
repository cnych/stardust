package csvx

import (
	"encoding/csv"
	"io"
	"os"
)

type Handler func(line int, record interface{}) (bool, error)

type Options struct {
	Comma            rune
	Comment          rune
	FieldsPerRecord  int
	LazyQuotes       bool
	TrailingComma    bool
	TrimLeadingSpace bool
	IgnoreErr        bool
}

func newReader(r io.Reader, opts *Options) *csv.Reader {
	reader := csv.NewReader(r)
	if opts.Comma != 0 {
		reader.Comma = opts.Comma
	}
	reader.Comment = opts.Comment
	reader.FieldsPerRecord = opts.FieldsPerRecord
	reader.LazyQuotes = opts.LazyQuotes
	reader.TrailingComma = opts.TrailingComma
	reader.TrimLeadingSpace = opts.TrimLeadingSpace
	return reader
}

func EachRecord(r io.Reader, h Handler, opts Options) error {
	reader := newReader(r, &opts)
	line := -1
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		line += 1
		if err != nil {
			if opts.IgnoreErr {
				continue
			} else {
				return err
			}
		}
		next, err := h(line, record)
		if err != nil {
			if opts.IgnoreErr {
				continue
			} else {
				return err
			}
		}
		if !next {
			break
		}
	}
	return nil
}

func FileEachRecord(filename string, h Handler, opts Options) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	return EachRecord(f, h, opts)
}
