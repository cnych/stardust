package config

import (
	"errors"
	"time"
)

type Holder interface {
	Config() interface{}
	Close() error
}

type ReloadableHolder interface {
	Holder
	ReloadListener() chan<- interface{}
	SetReloadListener(l chan<- interface{})
}

type Decoder func(data []byte) (interface{}, error)

var (
	emptyLocErr             error = errors.New("Location is empty")
	nilDecoderErr           error = errors.New("Decoder is nil")
	illegalCheckIntervalErr error = errors.New("CheckInterval must > 0")
)

func Load(loc string, decoder Decoder) (Holder, error) {
	if loc == "" {
		return nil, emptyLocErr
	}
	if decoder == nil {
		return nil, nilDecoderErr
	}

	data, err := loadData(loc)
	if err != nil {
		return nil, err
	}
	conf, err := decodeConfig(data, decoder)
	if err != nil {
		return nil, err
	}
	return &StaticHolder{conf}, nil
}

func MustLoad(loc string, decoder Decoder) Holder {
	h, err := Load(loc, decoder)
	if err != nil {
		panic(err)
	}
	return h
}

func AutoReload(loc string, checkInterval time.Duration, decoder Decoder) (ReloadableHolder, error) {
	if loc == "" {
		return nil, emptyLocErr
	}
	if decoder == nil {
		return nil, nilDecoderErr
	}

	newDefaultHandler := func() (ReloadableHolder, error) {
		h := &periodicHolder{loc: loc, checkInternal: checkInterval, decoder: decoder}
		err := h.start()
		if err != nil {
			return nil, err
		}
		return h, nil
	}

	if checkInterval <= 0 {
		return nil, illegalCheckIntervalErr
	}
	return newDefaultHandler()

}

func MustAutoReload(loc string, checkInternal time.Duration, decoder Decoder) ReloadableHolder {
	h, err := AutoReload(loc, checkInternal, decoder)
	if err != nil {
		panic(err)
	}
	return h
}
