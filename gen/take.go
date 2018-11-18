package gen

import "context"

// Take reads `n` values from the supplied value stream. By itself this may not seem very useful, but when combined with
// other generators this can lead to some clear, easy to work with code (see the example).
//
// Supports cancellation via the supplied context.
func Take(ctx context.Context, valStream <-chan interface{}, n int) <-chan interface{} {
	stream := make(chan interface{})

	go func() {
		defer close(stream)

		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case stream <- <-valStream:
			}
		}
	}()

	return stream
}
