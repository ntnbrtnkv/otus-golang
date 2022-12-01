package db

import (
	"github.com/jmoiron/sqlx"
	. "github.com/ntnbrtnkv/otus-golang/calendar/model"
	"go.uber.org/zap"
)

type CalendarDBStorage struct {
	db     *sqlx.DB
	logger *zap.SugaredLogger
}

func NewCalendarDBStorage(db *sqlx.DB, logger *zap.SugaredLogger) *CalendarDBStorage {
	return &CalendarDBStorage{db, logger}
}

func (m *CalendarDBStorage) AddEvent(event Event) error {
	q := "insert into Event (Title, Description, TimeStart, TimeEnd) values ($1, $2, $3, $4)"
	_, err := m.db.Exec(q, event.Title, event.Description, event.TimeStart, event.TimeEnd)
	return err
}

func (m *CalendarDBStorage) RemoveEvent(id string) error {
	q := "delete from Event where id = $1"
	_, err := m.db.Exec(q, id)
	return err
}

func (m *CalendarDBStorage) ChangeEvent(event Event) error {
	q := "update Event set Title = $1, Description = $2, TimeStart = $3, TimeEnd = $4 where ID = $5"
	_, err := m.db.Exec(q, event.Title, event.Description, event.TimeStart, event.TimeEnd, event.ID)
	return err
}

func (m *CalendarDBStorage) ListEvents() []Event {
	var res []Event
	q := "select ID, Title, Description, TimeStart, TimeEnd from Event"
	err := m.db.Select(&res, q)

	if err != nil {
		m.logger.Error(err)
		return res
	}

	return res
}
