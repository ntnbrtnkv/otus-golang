package model

import "time"

type Event struct {
	ID          string
	TimeStart   time.Time
	TimeEnd     time.Time
	Title       string
	Description string
}

func (e *Event) IsOverlapping(t Event) bool {
	return e.TimeEnd.After(t.TimeStart) && t.TimeEnd.After(e.TimeStart)
}
