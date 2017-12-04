package cache

import (
	"time"
)

type Put struct {
	Cache    Cache
	TTL      time.Duration
	MarkMiss bool
	MissTTL  time.Duration
}

type ValueCreator func(k string) (interface{}, error)

func Get(p *Put, k string, vc ValueCreator) (interface{}, error) {
	if p == nil || p.Cache == nil {
		return vc(k)
	}
	v, err := p.Cache.Get(k)
	if err != nil {
		return nil, err
	}

	if v != nil {
		if v == TheMissing {
			return nil, nil
		}
		return v, nil
	}
	v, err = vc(k)
	// 如果仅仅是mgo没有读到数据，则今天missing机制
	if err != nil && !IsNotFoundErr(err) {
		return nil, err
	}
	if v != nil {
		err = p.Cache.Put(k, v, p.TTL)
		if err != nil {
			return nil, err
		}
		return v, nil
	}
	if p.MarkMiss {
		p.Cache.MarkMissing(k, p.MissTTL)
	}
	return nil, nil

}

// 二级缓存
func Get2(p1, p2 *Put, k string, vm ValueCreator) (interface{}, error) {
	c1Nil, c2Nil := p1 == nil || p1.Cache == nil, p2 == nil || p2.Cache == nil

	if c1Nil && c2Nil {
		return vm(k)
	}
	if c1Nil {
		return Get(p2, k, vm)
	}
	if c2Nil {
		return Get(p1, k, vm)
	}
	return Get(p1, k, func(k string) (interface{}, error) {
		return Get(p2, k, vm)
	})
}
