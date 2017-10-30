package config

import "github.com/cnych/starjazz/logx"

func decodeConfig(data []byte, decoder Decoder) (interface{}, error) {
	c, err := decoder(data)
	if err != nil {
		logx.WithError(err).Error("config: Parse configuration data error")
	}
	return c, err
}
