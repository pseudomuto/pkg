package gen

import "context"

// Repeat repeats the supplied values to the returned channel
//
// Supports cancellation via the supplied context.
func Repeat(ctx context.Context, values ...interface{}) <-chan interface{} {
	stream := make(chan interface{})

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
func RepeatFunc(ctx context.Context, fn func() interface{}) <-chan interface{} {
	stream := make(chan interface{})

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
