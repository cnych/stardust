package useragent

import (
	"errors"
	"github.com/mssola/user_agent"
)

var (
	ErrParse = errors.New("Parse UA error")
)

type UA struct {
	UA                   string
	Mozilla              string
	Platform             string
	OS                   string
	Loc                  string
	Bot                  bool
	Mobile               bool
	BrowserEngine        string
	BrowserEngineVersion string
	BrowserName          string
	BrowserVersion       string
}

func Parse(s string, to *UA) (*UA, error) {
	ua0 := user_agent.New(s)
	if ua0 == nil {
		return nil, ErrParse
	}
	if to == nil {
		to = &UA{}
	}
	to.UA = ua0.UA()
	to.Mozilla = ua0.Mozilla()
	to.Platform = ua0.Platform()
	to.OS = ua0.OS()
	to.Loc = ua0.Localization()
	to.Bot = ua0.Bot()
	to.Mobile = ua0.Mobile()
	to.BrowserEngine, to.BrowserEngineVersion = ua0.Engine()
	to.BrowserName, to.BrowserVersion = ua0.Browser()
	return to, nil
}
