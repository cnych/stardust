package tomldec

import (
	"github.com/BurntSushi/toml"

	"github.com/cnych/starjazz/config"
)

func To(v interface{}) config.Decoder {
	return func(data []byte) (interface{}, error) {
		err := toml.Unmarshal(data, v)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
}
