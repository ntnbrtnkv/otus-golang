package inmemory

import (
	. "github.com/ntnbrtnkv/otus-golang/calendar/model"
	"sort"
)

type InMemoryStorage struct {
	events map[string]Event
}

func (s *InMemoryStorage) AddEvent(event Event) error {
	if s.events == nil {
		s.events = make(map[string]Event)
	}

	_, ok := s.events[event.ID]
	if ok {
		return DuplicateIdError
	}

	for _, ev := range s.events {
		if ev.IsOverlapping(event) {
			return BusyError
		}
	}

	s.events[event.ID] = event

	return nil
}

func (s *InMemoryStorage) RemoveEvent(id string) error {
	if s.events == nil {
		s.events = make(map[string]Event)
	}

	if _, ok := s.events[id]; !ok {
		return NotFoundError
	}

	delete(s.events, id)

	return nil
}

func (s *InMemoryStorage) ChangeEvent(event Event) error {
	if s.events == nil {
		s.events = make(map[string]Event)
	}

	if _, ok := s.events[event.ID]; !ok {
		return NotFoundError
	}

	s.events[event.ID] = event

	return nil
}

func (s *InMemoryStorage) ListEvents() []Event {
	var list []Event

	for _, v := range s.events {
		list = append(list, v)
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].TimeStart.Before(list[j].TimeStart)
	})

	return list
}
