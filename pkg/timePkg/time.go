package timepkg

import "time"

func Now() time.Time {
	return time.Now().UTC()
}

func NowUnixMilli() int64 {
	return time.Now().UTC().UnixMilli()
}

func ToUnixMilli(t time.Time) int64 {
	return t.UnixMilli()
}

func ToUnix(t time.Time) int64 {
	return t.Unix()
}

func AddDate(years int, months int, days int) time.Time {
	return time.Now().UTC().AddDate(years, months, days)
}

func Sleep(d time.Duration) {
	time.Sleep(d)
}

func SecondMultiply(d time.Duration) time.Duration {
	return time.Second * d
}
