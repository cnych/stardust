package syncx

import (
	"errors"
	"time"
)

type DistributedLock interface {
	TryDo(key string, ttl time.Duration, f func()) (bool, error)
}

func PrefixedDistributedLock(l DistributedLock, keyPrefix string) DistributedLock {
	if l == nil {
		panic(errors.New("The l is nil"))
	}
	return &prefixedDistributedLock{l, keyPrefix}
}

type prefixedDistributedLock struct {
	l         DistributedLock
	KeyPrefix string
}

func (l *prefixedDistributedLock) TryDo(key string, ttl time.Duration, f func()) (bool, error) {
	return l.l.TryDo(l.KeyPrefix+key, ttl, f)
}
