package timex

import (
	"time"
)

func After1H() time.Time { // 1小时后
	return time.Now().Add(DurationH(1))
}

func After1W() time.Time {
	return time.Now().Add(DurationH(1 * 24 * 7)) // 1周后
}

func After1Y() time.Time {
	return time.Now().Add(DurationH(1 * 24 * 360)) // 1年
}

func After10Y() time.Time {
	return time.Now().Add(DurationH(1 * 24 * 360 * 10)) // 10年
}

func DurationH(h int64) time.Duration {
	return time.Duration(int64(time.Hour) * h)
}

func DurationM(m int64) time.Duration {
	return time.Duration(int64(time.Minute) * m)
}

func DurationS(s int64) time.Duration {
	return time.Duration(int64(time.Second) * s)
}

func DurationMS(ms int64) time.Duration {
	return time.Duration(int64(time.Millisecond) * ms)
}
