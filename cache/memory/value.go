package memory

import "time"

type value struct {
	ExpireAt time.Time
	V        interface{}
}

func getValue(v0 *value, now time.Time) interface{} {
	if v0 == nil {
		return nil
	}
	if v0.ExpireAt.IsZero() {
		return v0.V
	} else {
		if now.After(v0.ExpireAt) { // 现在时间在expireAt之后，表示过期了
			return nil
		} else {
			return v0.V
		}
	}
}

func makeValue(v interface{}, ttl time.Duration, now time.Time) *value {
	v0 := &value{time.Time{}, v}
	if ttl > 0 {
		v0.ExpireAt = now.Add(ttl)
	}
	return v0
}
