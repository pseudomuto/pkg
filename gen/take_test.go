package gen_test

import (
	"context"
	"fmt"
	"testing"

	. "github.com/pseudomuto/pkg/gen"
	"github.com/stretchr/testify/assert"
)

func TestTake(t *testing.T) {
	inputStream := make(chan interface{}, 3)
	inputStream <- 1
	inputStream <- 2
	inputStream <- 3
	defer close(inputStream)

	i := 1
	for res := range Take(context.Background(), inputStream, 3) {
		func(v int) { assert.Equal(t, v, res.(int)) }(i)
		i++
	}

	t.Run("with cancellation", func(t *testing.T) {
		inputStream := make(chan interface{}, 2)
		inputStream <- 1
		inputStream <- 2
		defer close(inputStream)

		ctx, cancel := context.WithCancel(context.Background())
		cancel() // cancel immediately

		count := 0
		for range Take(ctx, inputStream, 2) {
			count++
		}

		assert.InDelta(t, count, 1, 1)
	})
}

// Print the number 1 ten times
func ExampleTake() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := range Take(ctx, Repeat(ctx, 1), 10) {
		fmt.Print(i)
	}

	// Output: 1111111111
}
