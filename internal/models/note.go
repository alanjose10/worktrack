package models

import "time"

type Note struct {
	ID      int
	TaskID  int
	Content string
	AddedOn time.Time
}
