package orchestrator

import (
	"fmt"
	"time"
)

type execFn func() error

// ExecuteWithTimeout - executes a function with a timeout
func ExecuteWithTimeout(fn execFn, name string, timeOutSeconds int) error {
	res := make(chan error, 1)
	go func() {
		res <- fn()
	}()

	select {
	case <-time.After(time.Duration(timeOutSeconds) * time.Second):
		return fmt.Errorf("timeout occurred after %d seconds, while executing function %q", timeOutSeconds, name)
	case result := <-res:
		return result
	}
}
