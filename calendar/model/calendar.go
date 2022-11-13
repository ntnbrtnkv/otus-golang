package model

type Calendar interface {
	AddEvent(event Event) error
	RemoveEvent(id string) error
	ChangeEvent(event Event) error
	ListEvents() []Event
}
