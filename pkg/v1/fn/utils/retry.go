package utils

import "time"

func retry[T any](retries int, execFn func() (*T, error)) (*T, error) {
	var response = new(T)
	var err error
	for i := 0; i < retries; i++ {
		response, err = execFn()
		if err != nil {
			time.Sleep(1 * time.Second)
			continue
		}
		break
	}

	return response, err
}
