package tasks

import (
	"fmt"
)

type TaskFunction func() error

type LimitReachedError struct {
	limit int
}

func (e LimitReachedError) Error() string {
	return fmt.Sprintf("error limit reached %d", e.limit)
}

func Worker(id int, tasks <-chan TaskFunction, out chan<- error) {
	fmt.Printf("Worker %d started\n", id)
	for f := range tasks {
		fmt.Printf("Worker %d starting task %p\n", id, f)
		out <- f()
		fmt.Printf("Worker %d done with task %p\n", id, f)
	}
	fmt.Printf("Worker %d stopped\n", id)
}

func Run(funcs []TaskFunction, maxWorkers int, maxFails int) error {
	tasks := make(chan TaskFunction, len(funcs))
	out := make(chan error)
	failsLeft := maxFails
	tasksNum := len(funcs)

	for _, f := range funcs {
		tasks <- f
	}

	wLen := tasksNum
	if wLen > maxWorkers {
		wLen = maxWorkers
	}

	for i := 0; i < wLen; i++ {
		go Worker(i, tasks, out)
	}

	close(tasks)

	for o := range out {
		tasksNum--
		fmt.Printf("Got result: %v, tasks left %d\n", o, tasksNum)
		if o != nil {
			failsLeft--
			if failsLeft == 0 {
				return LimitReachedError{
					limit: maxFails,
				}
			}
		}
		if tasksNum == 0 {
			close(out)
		}
	}

	return nil
}
