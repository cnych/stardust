package redisc

import (
	"errors"
	"github.com/cnych/starjazz/errorsx"
	"time"
)

func (c *Client) Get(k string) (interface{}, error) {
	data, err := c.c.Get(k).Bytes()

	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, err
	}
	return c.encoder.Decode(data)
}

func (c *Client) Get2(k string) (interface{}, error) {
	data, err := c.c.Get(k).Bytes()

	if err != nil {
		if err.Error() == "redis: nil" {
			return nil, nil
		}
		return nil, err
	}
	return c.encoder.Decode(data)
}

func (c *Client) Put(k string, v interface{}, ttl time.Duration) error {
	data, err := c.encoder.Encode(v)
	if err != nil {
		return err
	}
	_, err = c.c.Set(k, data, ttl).Result()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) MarkMissing(k string, ttl time.Duration) error {
	_, err := c.c.Set(k, theMissingBytes, ttl).Result()
	return err
}

func (c *Client) Remove(k string) error {
	_, err := c.c.Del(k).Result()
	return err
}

func (c *Client) MGet(ks []string) (map[string]interface{}, error) {
	r := make(map[string]interface{})
	if len(ks) == 0 {
		return r, nil
	}
	cr, err := c.c.MGet(ks...).Result()
	if err != nil {
		return nil, err
	}
	if len(cr) != len(ks) {
		return nil, errors.New("MGet len(cr) != len(ks)")
	}
	for i, k := range ks {
		v00 := cr[i]
		v0, ok := v00.([]byte)
		if !ok {
			return nil, errors.New("The value is not []byte")
		}
		if !isMissing(v0) {
			r[k] = nil
		} else {
			v, err := c.encoder.Decode(v0)
			if err != nil {
				return nil, err
			}
			r[k] = v
		}
	}
	return r, nil
}

func (c *Client) MPut(kvs map[string]interface{}, ttl time.Duration) error {
	if len(kvs) == 0 {
		return nil
	}

	if ttl > 0 {
		var errs []error
		for k, v := range kvs {
			err0 := c.Put(k, v, ttl)
			errs = append(errs, err0)
		}
		return errorsx.Multi(errs)
	} else {
		pairs := make([]interface{}, 0, len(kvs)*2)
		for k, v := range kvs {
			v1, err := c.encoder.Encode(v)
			if err != nil {
				return err
			}
			pairs = append(pairs, k, v1)
		}
		_, err := c.c.MSet(pairs).Result()
		return err
	}
}

func (c *Client) MMarkMissing(ks []string, ttl time.Duration) error {
	if len(ks) == 0 {
		return nil
	}
	if ttl > 0 {
		var errs []error
		for _, k := range ks {
			err0 := c.MarkMissing(k, ttl)
			errs = append(errs, err0)
		}
		return errorsx.Multi(errs)
	} else {
		pairs := make([]interface{}, 0, len(ks)*2)
		for _, k := range ks {
			pairs = append(pairs, k, theMissingBytes)
		}
		_, err := c.c.MSet(pairs).Result()
		return err
	}
	return nil
}

func (c *Client) MRemove(ks []string) error {
	if len(ks) == 0 {
		return nil
	}
	_, err := c.c.Del(ks...).Result()
	return err
}
