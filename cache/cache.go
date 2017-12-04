package cache

import "time"

var (
	TheMissing interface{} = "<<_MISSING_>>"
)

// 抽象缓存接口
type Cache interface {
	// Single
	Get(k string) (interface{}, error)
	Put(k string, v interface{}, ttl time.Duration) error
	MarkMissing(k string, ttl time.Duration) error
	Remove(k string) error

	// Multi
	MGet(ks []string) (map[string]interface{}, error)
	MPut(kvs map[string]interface{}, ttl time.Duration) error
	MMarkMissing(ks []string, ttl time.Duration) error
	MRemove(ks []string) error
}
