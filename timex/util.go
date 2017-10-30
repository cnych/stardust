package timex

import (
	"time"
)

func NowSecT() time.Time {
	return time.Now().Truncate(time.Second)
}
