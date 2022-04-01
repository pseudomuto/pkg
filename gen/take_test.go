package gen_test

import (
	"context"
	"testing"

	. "github.com/pseudomuto/pkg/gen"
	"github.com/stretchr/testify/require"
)

func TestTake(t *testing.T) {
	inputStream := make(chan interface{}, 3)
	inputStream <- 1
	inputStream <- 2
	inputStream <- 3
	defer close(inputStream)

	i := 1
	for res := range Take(context.Background(), inputStream, 3) {
		func(v int) { require.Equal(t, v, res) }(i)
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

		require.InDelta(t, count, 1, 1)
	})
}
