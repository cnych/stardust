package timex

import (
	"errors"
	"fmt"
	"time"
)

var (
	ErrParse = errors.New("Parse time error")
)

var (
	timeLayoutsForParse = []string{
		"20060102150405",
		time.ANSIC,
		time.UnixDate,
		time.RubyDate,
		time.Kitchen,
		time.RFC3339,
		time.RFC3339Nano,
		"20060102",
		"2006-01-02",                         // RFC 3339
		"2006-01-02 15:04",                   // RFC 3339 with minutes
		"2006-01-02 15:04:05",                // RFC 3339 with seconds
		"2006-01-02 15:04:05-07:00",          // RFC 3339 with seconds and timezone
		"2006-01-02T15Z0700",                 // ISO8601 with hour
		"2006-01-02T15:04Z0700",              // ISO8601 with minutes
		"2006-01-02T15:04:05Z0700",           // ISO8601 with seconds
		"2006-01-02T15:04:05.999999999Z0700", // ISO8601 with nanoseconds
	}
)

func TryParse(s string) (time.Time, error) {
	for _, layout := range timeLayoutsForParse {
		r, err := time.Parse(layout, s)
		if err == nil {
			return r, nil
		}
	}
	return time.Time{}, ErrParse
}

func MustParse(s string) time.Time {
	t, err := TryParse(s)
	if err != nil {
		panic(err)
	}
	return t
}

// returns a formatted string of `time.RFC1123` format.
func TimeStrToRFC1123(str string) string {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		t, err = time.Parse(time.RFC1123, str)
		if err != nil {
			panic("Time format invalid. The time format must be time.RFC3339 or time.RFC1123")
		}
	}
	return t.Format(time.RFC1123)
}

// returns a utc string of a time instance.
func TimeToUTCStr(t time.Time) string {
	format := time.RFC3339 // 2006-01-02T15:04:05Z07:00
	return t.UTC().Format(format)
}

func TimestampToHourStr(ts int) string {
	//格式化为字符串,tm为Time类型
	tm := time.Unix(int64(ts), 0)
	return tm.Format("3小时4分5秒")
}

func GetISO8601TimeStamp(ts time.Time) string {
	t := ts.UTC()
	cc := fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02dZ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
	fmt.Printf("#######################: %s\n", cc)
	return fmt.Sprintf("%04d-%02d-%02dT%02d:%02d:%02dZ", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}
