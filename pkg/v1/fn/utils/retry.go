package utils

import "time"

// Retry - will retry the function if errors out
func Retry[T any](retries int, retryInterval time.Duration, execFn func() (T, error)) (T, error) {
	var response T
	var err error

	for i := 0; i < retries; i++ {
		response, err = execFn()
		if err != nil {
			time.Sleep(retryInterval)
			continue
		}
		break
	}

	return response, err
}
