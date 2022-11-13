package model

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestIsOverlapping(t *testing.T) {
	e := Event{TimeStart: time.UnixMicro(5), TimeEnd: time.UnixMicro(10)}

	require.True(t, e.IsOverlapping(Event{TimeStart: time.UnixMicro(4), TimeEnd: time.UnixMicro(11)}))
	require.True(t, e.IsOverlapping(Event{TimeStart: time.UnixMicro(4), TimeEnd: time.UnixMicro(6)}))
	require.True(t, e.IsOverlapping(Event{TimeStart: time.UnixMicro(9), TimeEnd: time.UnixMicro(11)}))

	require.False(t, e.IsOverlapping(Event{TimeStart: time.UnixMicro(10), TimeEnd: time.UnixMicro(11)}))
	require.False(t, e.IsOverlapping(Event{TimeStart: time.UnixMicro(3), TimeEnd: time.UnixMicro(4)}))
}
