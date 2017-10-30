package config

import (
	"bytes"
	"github.com/cnych/starjazz/logx"
	"github.com/cnych/starjazz/syncx"
	"sync"
	"time"
)

type periodicHolder struct {
	loc            string
	decoder        Decoder
	checkInternal  time.Duration
	mutex          sync.RWMutex
	lastConfig     interface{}
	lastData       []byte
	closeChan      chan int
	reloadListener chan<- interface{}
}

func (h *periodicHolder) start() error {
	data, err := loadData(h.loc)
	if err != nil {
		return err
	}
	conf, err := decodeConfig(data, h.decoder)
	if err != nil {
		logx.WithError(err).WithField("loc", h.loc).Error("Load error")
		return err
	}
	syncx.WLockDo(&h.mutex, func() {
		h.lastData, h.lastConfig = data, conf
		logx.WithField("loc", h.loc).Debug("Load OK")
	})
	h.closeChan = make(chan int)
	go func(closeChan <-chan int) {
	STOP:
		for {
			select {
			case _, ok := <-closeChan:
				if !ok {
					break STOP
				}
			case <-time.After(h.checkInternal):
				{
					data, err := loadData(h.loc)
					if err != nil {
						continue
					}
					syncx.WLockDo(&h.mutex, func() {
						if bytes.Equal(data, h.lastData) {
							return
						}
						conf, err := h.decoder(data)
						if err != nil {
							logx.WithError(err).WithField("loc", h.loc).Error("Reload error")
							return
						}
						h.lastData, h.lastConfig = data, conf
						logx.WithField("loc", h.loc).Debug("Reload OK")
						if l := h.reloadListener; l != nil {
							l <- conf
						}
					})
				}
			}
		}
	}(h.closeChan)
	return nil
}

func (h *periodicHolder) Config() interface{} {
	var r interface{} = nil
	syncx.RLockDo(&h.mutex, func() {
		r = h.lastConfig
	})
	return r
}

func (h *periodicHolder) Close() error {
	close(h.closeChan)
	h.closeChan = nil
	return nil
}

func (h *periodicHolder) ReloadListener() chan<- interface{} {
	return h.reloadListener
}

func (h *periodicHolder) SetReloadListener(l chan<- interface{}) {
	h.reloadListener = l
}
