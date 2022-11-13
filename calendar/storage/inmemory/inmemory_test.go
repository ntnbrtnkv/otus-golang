package inmemory

import (
	"testing"
	"time"

	. "github.com/ntnbrtnkv/otus-golang/calendar/model"
	"github.com/stretchr/testify/require"
)

func TestAddEvent(t *testing.T) {
	cal := InMemoryStorage{}

	testEv := Event{
		ID:          "1",
		TimeStart:   time.Unix(4, 0),
		TimeEnd:     time.Unix(10, 0),
		Title:       "Test",
		Description: "Test it",
	}

	runEv := Event{
		ID:          "2",
		TimeStart:   time.Unix(1, 0),
		TimeEnd:     time.Unix(4, 0),
		Title:       "Run",
		Description: "Run it",
	}

	cal.AddEvent(testEv)
	cal.AddEvent(runEv)

	require.Equal(t, map[string]Event{"2": runEv, "1": testEv}, cal.events, "Sort by TimeStart")
}

func TestAddEventWithSameID(t *testing.T) {
	cal := InMemoryStorage{}

	testEv := Event{
		ID:          "1",
		TimeStart:   time.Unix(4, 0),
		TimeEnd:     time.Unix(10, 0),
		Title:       "Test",
		Description: "Test it",
	}

	runEv := Event{
		ID:          "1",
		TimeStart:   time.Unix(1, 0),
		TimeEnd:     time.Unix(5, 0),
		Title:       "Run",
		Description: "Run it",
	}

	cal.AddEvent(testEv)
	err := cal.AddEvent(runEv)

	require.Equal(t, map[string]Event{"1": testEv}, cal.events)
	require.Equal(t, DuplicateIdError, err)
}

func TestAddEventOverlapping(t *testing.T) {
	cal := InMemoryStorage{}

	testEv := Event{
		ID:          "1",
		TimeStart:   time.Unix(4, 0),
		TimeEnd:     time.Unix(10, 0),
		Title:       "Test",
		Description: "Test it",
	}

	runEv := Event{
		ID:          "2",
		TimeStart:   time.Unix(1, 0),
		TimeEnd:     time.Unix(5, 0),
		Title:       "Run",
		Description: "Run it",
	}

	cal.AddEvent(testEv)
	err := cal.AddEvent(runEv)

	require.Equal(t, map[string]Event{"1": testEv}, cal.events)
	require.Equal(t, BusyError, err)
}

func TestRemoveEvent(t *testing.T) {
	cal := InMemoryStorage{}

	testEv := Event{
		ID:          "1",
		TimeStart:   time.Unix(4, 0),
		TimeEnd:     time.Unix(10, 0),
		Title:       "Test",
		Description: "Test it",
	}

	runEv := Event{
		ID:          "2",
		TimeStart:   time.Unix(1, 0),
		TimeEnd:     time.Unix(4, 0),
		Title:       "Run",
		Description: "Run it",
	}

	cal.AddEvent(testEv)
	cal.AddEvent(runEv)
	err := cal.RemoveEvent(runEv.ID)

	require.Equal(t, map[string]Event{"1": testEv}, cal.events)
	require.Nil(t, err)
}

func TestRemoveEventNonExisting(t *testing.T) {
	cal := InMemoryStorage{}

	testEv := Event{
		ID:          "1",
		TimeStart:   time.Unix(4, 0),
		TimeEnd:     time.Unix(10, 0),
		Title:       "Test",
		Description: "Test it",
	}

	runEv := Event{
		ID:          "2",
		TimeStart:   time.Unix(1, 0),
		TimeEnd:     time.Unix(4, 0),
		Title:       "Run",
		Description: "Run it",
	}

	err := cal.RemoveEvent("3")

	require.Equal(t, map[string]Event{}, cal.events)
	require.Equal(t, NotFoundError, err)

	cal.AddEvent(testEv)
	cal.AddEvent(runEv)
	err = cal.RemoveEvent("3")

	require.Equal(t, map[string]Event{"2": runEv, "1": testEv}, cal.events)
	require.Equal(t, NotFoundError, err)
}

func TestChangeEvent(t *testing.T) {
	cal := InMemoryStorage{}

	testEv := Event{
		ID:          "1",
		TimeStart:   time.Unix(4, 0),
		TimeEnd:     time.Unix(10, 0),
		Title:       "Test",
		Description: "Test it",
	}

	nextEv := Event{
		ID:          "1",
		TimeStart:   time.Unix(4, 0),
		TimeEnd:     time.Unix(10, 0),
		Title:       "Test 2",
		Description: "Test it",
	}

	cal.AddEvent(testEv)
	err := cal.ChangeEvent(nextEv)

	require.Equal(t, map[string]Event{"1": nextEv}, cal.events)
	require.Nil(t, err)
}

func TestChangeEventNonExisting(t *testing.T) {
	cal := InMemoryStorage{}

	testEv := Event{
		ID:          "1",
		TimeStart:   time.Unix(4, 0),
		TimeEnd:     time.Unix(10, 0),
		Title:       "Test",
		Description: "Test it",
	}

	nextEv := Event{
		ID:          "2",
		TimeStart:   time.Unix(4, 0),
		TimeEnd:     time.Unix(10, 0),
		Title:       "Test 2",
		Description: "Test it",
	}

	err := cal.ChangeEvent(nextEv)

	require.Equal(t, map[string]Event{}, cal.events)
	require.Equal(t, NotFoundError, err)

	cal.AddEvent(testEv)
	err = cal.ChangeEvent(nextEv)

	require.Equal(t, map[string]Event{"1": testEv}, cal.events)
	require.Equal(t, NotFoundError, err)
}

func TestList(t *testing.T) {
	cal := InMemoryStorage{}

	testEv := Event{
		ID:          "1",
		TimeStart:   time.Unix(4, 0),
		TimeEnd:     time.Unix(10, 0),
		Title:       "Test",
		Description: "Test it",
	}

	runEv := Event{
		ID:          "2",
		TimeStart:   time.Unix(1, 0),
		TimeEnd:     time.Unix(4, 0),
		Title:       "Run",
		Description: "Run it",
	}

	cal.AddEvent(testEv)
	cal.AddEvent(runEv)

	require.Equal(t, []Event{runEv, testEv}, cal.ListEvents())
}
