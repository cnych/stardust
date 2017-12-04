package config

import (
	"fmt"
	"github.com/franela/goreq"
	"io/ioutil"
	"net/http"
	"strings"
)

// 加载toml数据，支持网络地址
func loadData(loc string) ([]byte, error) {
	if strings.HasPrefix(loc, "http://") || strings.HasPrefix(loc, "https://") {
		return loadHTTP(loc)
	}
	return loadFile(loc)
}

func loadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func loadHTTP(urlStr string) ([]byte, error) {
	resp, err := goreq.Request{Method: http.MethodGet, Uri: urlStr}.Do()
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response status error (%d)", resp.StatusCode)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
