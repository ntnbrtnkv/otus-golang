package tasks

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestRunAllPositive(t *testing.T) {
	tasks := []TaskFunction{
		func() error {
			time.Sleep(time.Second)
			return nil
		},
		func() error {
			time.Sleep(2000 * time.Millisecond)
			return nil
		},
	}

	require.Equal(t, nil, Run(tasks, 2, 1))
}

func TestRunReachErrorsLimit(t *testing.T) {
	i := 0
	tasks := []TaskFunction{
		func() error {
			time.Sleep(time.Second)
			i++
			return nil
		},
		func() error {
			time.Sleep(1300 * time.Millisecond)
			i++
			return fmt.Errorf("error 1")
		},
		func() error {
			time.Sleep(1600 * time.Millisecond)
			i++
			return fmt.Errorf("error 2")
		},
		func() error {
			time.Sleep(1500 * time.Millisecond)
			i++
			return nil
		},
	}

	err := Run(tasks, 2, 1)
	err, ok := err.(LimitReachedError)
	require.True(t, ok)
	require.GreaterOrEqual(t, 3, i)
}

func TestRunReachErrorsLimitAtTheEnd(t *testing.T) {
	i := 0
	tasks := []TaskFunction{
		func() error {
			time.Sleep(time.Second)
			i++
			return nil
		},
		func() error {
			time.Sleep(1300 * time.Millisecond)
			i++
			return fmt.Errorf("error 1")
		},
		func() error {
			time.Sleep(1600 * time.Millisecond)
			i++
			return fmt.Errorf("error 2")
		},
		func() error {
			time.Sleep(1300 * time.Millisecond)
			i++
			return nil
		},
	}

	err := Run(tasks, 2, 2)
	err, ok := err.(LimitReachedError)
	require.Equal(t, true, ok)
	require.Equal(t, 4, i)
}
