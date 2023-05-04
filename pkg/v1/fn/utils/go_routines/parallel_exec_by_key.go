package go_routines

import (
	"context"
	"sync"
)

// ParallelExecByKey - fetches items in parallel, preserving the order of the keys.
// Takes a context, batch size, keys of type K and a function to instruct how to fetch the item by the key K.
// Returns []R or error
func ParallelExecByKey[R, K any](ctx context.Context, batchSize int, keys []K, fn func(key K) (*R, error)) ([]*R, error) {

	ctxCancel, cancelFunc := context.WithCancel(ctx)
	var buffer = make(chan bool, batchSize)
	var errCh = make(chan error, batchSize)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var syncValues = make([]*R, len(keys))

	for idx, key := range keys {
		wg.Add(1)
		go func(c context.Context, i int, k K) {
			defer wg.Done()
			buffer <- true
			defer func() { <-buffer }()

			select {
			case <-c.Done():
				return
			default:
				result, err := fn(k)
				if err != nil {
					errCh <- err
					cancelFunc()
					return
				}

				mu.Lock()
				syncValues[i] = result
				mu.Unlock()
			}
		}(ctxCancel, idx, key)
	}

	wg.Wait()
	cancelFunc()
	close(buffer)

	if len(errCh) > 0 {
		err := <-errCh
		close(errCh)
		return nil, err
	}

	return syncValues, nil
}
