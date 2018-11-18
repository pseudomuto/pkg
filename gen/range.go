package gen

import "context"

// Range yields the first `n` integers (one at a time). It's totally cool to use this with large numbers since the
// results are sent one at a time (not all stored in memory).
//
// Supports cancellation via the supplied context.
func Range(ctx context.Context, n int) <-chan int {
	return RangeFrom(ctx, 0, n)
}

// RangeFrom yields the first `n` integers after start. It yields the range [start, start+n).
func RangeFrom(ctx context.Context, start, n int) <-chan int {
	stream := make(chan int)

	go func() {
		defer close(stream)

		for i := 0; i < n; i++ {
			select {
			case <-ctx.Done():
				return
			case stream <- start + i:
			}
		}
	}()

	return stream
}
