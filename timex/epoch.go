package timex

import (
	"time"
)

func ToEpoch(t time.Time) int64 {
	if t.IsZero() {
		return 0
	}
	return t.Unix()
}

func ToEpochMS(t time.Time) int64 {
	return t.UnixNano() / int64(time.Millisecond)
}

func NowEpoch() int64 {
	return ToEpoch(time.Now())
}

func NowEpochMS() int64 {
	return ToEpochMS(time.Now())
}

func FromEpoch(sec int64) time.Time {
	if sec <= 0 {
		return time.Time{}
	}
	return time.Unix(sec, 0)
}

// TODO: FromEpochMS
