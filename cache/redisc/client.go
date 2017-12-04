package redisc

import (
	"bytes"
	"errors"
	"gopkg.in/redis.v4"
	"github.com/cnych/starjazz/cache"
	"github.com/cnych/starjazz/redisx"
)

var (
	theMissingBytes = []byte(cache.TheMissing.(string))
)

type Client struct {
	c       redis.Cmdable
	encoder Encoder
}

func New(addr redisx.Address, encoder Encoder) (*Client, error) {
	if encoder == nil {
		return nil, errors.New("Encoder is nil")
	}
	c, err := redisx.Dial(addr)
	if err != nil {
		return nil, err
	}
	return &Client{c: c, encoder: encoder}, nil
}

func isMissing(v []byte) bool {
	if v == nil {
		return false
	}
	return bytes.Equal(v, theMissingBytes)
}
