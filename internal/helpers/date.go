package helpers

import "time"

func GetHumanDate(t time.Time) string {
	return t.Format("2 Jan 2006 at 3:04 PM")
}
