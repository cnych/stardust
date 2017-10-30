package main

import (
	"fmt"
	"github.com/cnych/starjazz/config"
	"github.com/cnych/starjazz/config/tomldec"
	"github.com/cnych/starjazz/logx"
)

type Demo struct {
	Site string `toml:"site"`
}

type StableConfig struct {
	Demo `toml:"demo"`
}

func main() {
	dec := func(data []byte) (interface{}, error) {
		sci, err := tomldec.To(&StableConfig{})(data)
		if err != nil {
			return nil, err
		}
		// 设置默认配置
		sc := sci.(*StableConfig)
		return sc, nil
	}
	loc := "test.toml"
	scHolder, err := config.Load(loc, dec)
	if err != nil {
		logx.WithField("toml", loc).WithError(err).Error("Load Config toml")
	} else {
		sc := scHolder.Config().(*StableConfig)
		logx.WithField("toml", loc).WithField("result", fmt.Sprintf("%+v", sc.Demo)).Debug("Load Config toml")
	}
}
