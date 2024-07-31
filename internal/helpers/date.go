package helpers

import "time"

func GetHumanDate(t time.Time) string {
	return t.Format("2 Jan 2006 at 3:04 PM")
}

func GetTimeFromUnix(ts int64) time.Time {
	return time.Unix(ts, 0)
}

func GetCurrentDate() time.Time {
	return time.Now()
}

func ConvertToUnix(t time.Time) int64 {
	return t.Unix()
}

func GetYear(t time.Time) int {
	return t.Year()
}

func GetMonth(t time.Time) int {
	return int(t.Month())
}

func GetDay(t time.Time) int {
	return t.Day()
}
