package main

import (
	"fmt"
	"runtime"
)

func tasksRunner(tasks []func() error, goroutinesCount, errorsLimit int) error {
	if len(tasks) == 0 {
		return nil
	}

	tasksQueue := make(chan func() error, goroutinesCount)
	tasksResult := make(chan error, goroutinesCount)

	for i := 0; i < goroutinesCount; i++ {
		go func(tasksQueue <-chan func() error, tasksResult chan<- error) {
			for task := range tasksQueue {
				result := task()
				tasksResult <- result
				runtime.Gosched()
			}
		}(tasksQueue, tasksResult)
	}

	go func(tasks []func() error, tasksQueue chan<- func() error) {
		for _, task := range tasks {
			tasksQueue <- task
		}
	}(tasks, tasksQueue)

	errorsCount := 0
	tasksCount := 0

	for result := range tasksResult {
		tasksCount += 1
		if result != nil {
			errorsCount += 1
		}

		if errorsCount > errorsLimit {
			close(tasksQueue)
			close(tasksResult)
			return fmt.Errorf("errors limit in %d exceeded", errorsLimit)
		}
		if tasksCount == len(tasks) {
			close(tasksQueue)
			close(tasksResult)
			return nil
		}
	}
	return fmt.Errorf("something unexpected happened")
}
