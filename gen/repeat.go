package gen

import "context"

// Repeat repeats the supplied values to the returned channel.
//
// Supports cancellation via the supplied context.
func Repeat[T any](ctx context.Context, values ...T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)

		for {
			for _, v := range values {
				select {
				case <-ctx.Done():
					return
				case stream <- v:
				}
			}
		}
	}()

	return stream
}

// RepeatFunc calls the supplied function sending results to the returned channel.
//
// Supports cancellation via the supplied context.
func RepeatFunc[T any](ctx context.Context, fn func() T) <-chan T {
	stream := make(chan T)

	go func() {
		defer close(stream)
		for {
			select {
			case <-ctx.Done():
				return
			case stream <- fn():
			}
		}
	}()

	return stream
}
