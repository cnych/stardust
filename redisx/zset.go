package redisx

import "gopkg.in/redis.v4"

func ZAdd(c redis.Cmdable, key string, score float64, member string) (bool, error) {
	added, err := c.ZAdd(key, redis.Z{Score: score, Member: member}).Result()
	if err != nil {
		return false, err
	}
	if added == 1 {
		return true, nil
	}
	return false, nil
}

func zRevRangeWithScores(c redis.Cmdable, key string, star, stop int64) ([]redis.Z, error) {
	return c.ZRevRangeWithScores(key, star, stop).Result()
}

func ZAllRange(c redis.Cmdable, key string) ([]redis.Z, error) {
	return zRevRangeWithScores(c, key, 0, -1)
}

func ZAllCount(c redis.Cmdable, key string) (int64, error) {
	return c.ZCount(key, "-inf", "+inf").Result()
}

func ZRank(c redis.Cmdable, key string, member string) (int64, error) {
	return c.ZRank(key, member).Result()
}
