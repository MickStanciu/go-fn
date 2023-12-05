package batch

import (
	"context"
	"sync"
)

// ParallelExecByKey - given a set of keys, will batch execute `fn` for each key
// Takes a context, batch size, keys of type string and a function to instruct how to fetch the item by the key.
// Returns map[Key]R or error
func ParallelExecByKey[R any, Key string](ctx context.Context, batchSize int, keys []Key, fn func(ctx context.Context, key Key) (R, error)) (map[Key]R, error) {
	ctxCancel, cancelFunc := context.WithCancel(ctx)
	var buffer = make(chan bool, batchSize)
	var errCh = make(chan error, batchSize)

	var wg sync.WaitGroup
	var mu sync.Mutex
	var finalResult = make(map[Key]R)

	for _, key := range keys {
		wg.Add(1)
		go func(c context.Context, k Key) {
			defer wg.Done()
			buffer <- true
			defer func() { <-buffer }()

			select {
			case <-c.Done():
				return
			default:
				result, err := fn(ctxCancel, k)
				if err != nil {
					errCh <- err
					cancelFunc()
					return
				}

				mu.Lock()
				finalResult[k] = result
				mu.Unlock()
			}
		}(ctxCancel, key)
	}

	wg.Wait()
	cancelFunc()
	close(buffer)

	if len(errCh) > 0 {
		err := <-errCh
		close(errCh)
		return nil, err
	}

	return finalResult, nil
}
