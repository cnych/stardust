package redisx

import (
	"errors"
	"gopkg.in/redis.v4"
	"github.com/cnych/starjazz/timex"
	"time"
)

type Address struct {
	Addrs       []string `json:"addrs" toml:"addrs"`
	DB          int      `json:"db,omitempty" toml:"db"`
	Password    string   `json:"password,omitempty" toml:"password"`
	IdleTimeout int64    `json:"idle_timeout" toml:"idel_timeout"`
	MaxRetries  int      `json:"max_retries,omitempty" toml:"max_retries"`
	Cluster     bool     `json:"cluster,omitempty" toml:"cluster"`
}

func Dial(addr Address) (cmd redis.Cmdable, err error) {
	idleTS := timex.DurationS(int64(5))
	if addr.IdleTimeout > 0 {
		idleTS = timex.DurationS(addr.IdleTimeout)
	}
	switch len(addr.Addrs) {
	case 0:
		cmd = nil
		err = errors.New("Missing addrs")
	case 1:
		cmd = redis.NewClient(&redis.Options{
			Addr:        addr.Addrs[0],
			DB:          addr.DB,
			Password:    addr.Password,
			IdleTimeout: idleTS,
			MaxRetries:  addr.MaxRetries,
		})
		err = nil
	default:
		if addr.Cluster {
			cmd = redis.NewClusterClient(&redis.ClusterOptions{
				Addrs:       addr.Addrs,
				Password:    addr.Password,
				IdleTimeout: idleTS,
			})
		} else {
			addrMap := map[string]string{}
			for _, addr1 := range addr.Addrs {
				addrMap[addr1] = addr1
			}
			cmd = redis.NewRing(&redis.RingOptions{
				Addrs:       addrMap,
				DB:          addr.DB,
				Password:    addr.Password,
				MaxRetries:  addr.MaxRetries,
				IdleTimeout: idleTS,
			})
		}
		err = nil
	}
	go func() {
		for {
			cmd.Ping()
			time.Sleep(idleTS - 10)
		}
	}()
	return
}

func Close(r redis.Cmdable) error {
	type closeable interface {
		Close() error
	}
	c, ok := r.(closeable)
	if !ok {
		return errors.New("Not closable")
	}
	return c.Close()
}
