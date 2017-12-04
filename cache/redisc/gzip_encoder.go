package redisc

import (
	"compress/gzip"
	"github.com/cnych/starjazz/compressx/gzipx"
)

type GzipEncoder struct {
	Encoder Encoder
}

func (enc *GzipEncoder) Encode(v interface{}) ([]byte, error) {
	uncompressed, err := enc.Encoder.Encode(v)
	if err != nil {
		return nil, err
	}
	return gzipx.Compress(uncompressed, gzip.DefaultCompression)
}

func (enc *GzipEncoder) Decode(compressed []byte) (interface{}, error) {
	uncompressed, err := gzipx.Uncompress(compressed)
	if err != nil {
		return nil, err
	}
	return enc.Encoder.Decode(uncompressed)
}
