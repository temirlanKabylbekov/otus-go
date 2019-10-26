package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestPassEmptySetOfTasks(t *testing.T) {
	err := tasksRunner([]func() error{}, 2, 1)
	require.Nil(t, err)
}

func TestRunTasksInSingleGoroutine(t *testing.T) {
	tasks := []func() error{
		func() error {
			time.Sleep(1 * time.Second)
			return nil
		},
		func() error {
			time.Sleep(2 * time.Second)
			return nil
		},
		func() error {
			time.Sleep(1 * time.Second)
			return nil
		},
	}
	err := tasksRunner(tasks, 1, 1)
	require.Nil(t, err)
}

func TestErrorsLimitExceeded(t *testing.T) {
	tasks := []func() error{
		func() error {
			time.Sleep(1 * time.Second)
			return fmt.Errorf("task-1")
		},
		func() error {
			time.Sleep(2 * time.Second)
			return fmt.Errorf("task-2")
		},
		func() error {
			time.Sleep(1 * time.Second)
			return nil
		},
	}
	err := tasksRunner(tasks, 3, 1)
	require.EqualError(t, err, "errors limit in 1 exceeded")
}

func TestErrorsLimitNotExceeded(t *testing.T) {
	tasks := []func() error{
		func() error {
			time.Sleep(1 * time.Second)
			return fmt.Errorf("task-1")
		},
		func() error {
			time.Sleep(2 * time.Second)
			return nil
		},
		func() error {
			time.Sleep(1 * time.Second)
			return nil
		},
	}
	err := tasksRunner(tasks, 3, 2)
	require.Nil(t, err)
}

func TestRunSmallCountOfTasksInManyCoroutines(t *testing.T) {
	tasks := []func() error{
		func() error {
			time.Sleep(1 * time.Second)
			return nil
		},
		func() error {
			time.Sleep(2 * time.Second)
			return nil
		},
	}
	err := tasksRunner(tasks, 10, 1)
	require.Nil(t, err)
}
