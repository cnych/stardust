package gzipx

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io/ioutil"
)

var (
	errNilData          = errors.New("Data is nil")
	errUncompletedWrite = errors.New("Uncomplete write")
)

func Compress(data []byte, level int) ([]byte, error) {
	if data == nil {
		return nil, errNilData
	}

	buff := new(bytes.Buffer)
	w, err := gzip.NewWriterLevel(buff, level)
	if err != nil {
		return nil, err
	}
	_, err = w.Write(data)
	if err != nil {
		return nil, err
	}
	err = w.Close()
	if err != nil {
		return nil, err
	}
	return buff.Bytes(), nil
}

func Uncompress(data []byte) ([]byte, error) {
	if data == nil {
		return nil, errNilData
	}
	r, err := gzip.NewReader(bytes.NewReader(data))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	return ioutil.ReadAll(r)
}
