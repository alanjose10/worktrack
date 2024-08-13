package helpers

import "time"

func GetHumanDate(t time.Time) string {
	if t.IsZero() {
		return "NULL"
	}
	return t.Format("Mon 2 Jan 2006")
}

func GetTimeFromUnix(ts int64) time.Time {
	return time.Unix(ts, 0)
}

func GetCurrentDate() time.Time {
	return time.Now()
}

func GetYesterdayDate() time.Time {
	return GetCurrentDate().AddDate(0, 0, -1)
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

func ParseDate(date string) (time.Time, error) {
	return time.Parse("02-01-2006", date)
}
