package gen_test

import (
	"context"
	"math/rand"
	"testing"

	. "github.com/pseudomuto/pkg/gen"
	"github.com/stretchr/testify/require"
)

func TestRepeat(t *testing.T) {
	tests := []struct {
		inputs  []int
		outputs []int
	}{
		{
			inputs:  []int{1},
			outputs: []int{1, 1, 1, 1, 1},
		},
		{
			inputs:  []int{1, 2, 3},
			outputs: []int{1, 2, 3, 1, 2, 3, 1, 2, 3, 1},
		},
	}

	ctx := context.Background()
	for _, test := range tests {
		stream := Repeat(ctx, test.inputs...)

		for _, val := range test.outputs {
			require.Equal(t, val, <-stream)
		}
	}

	t.Run("with cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // cancel immediately

		count := 0
		for range Repeat(ctx, 1) {
			count++
		}

		require.InDelta(t, count, 1, 1)
	})
}

func TestRepeatFunc(t *testing.T) {
	rand := func() interface{} { return rand.Int() }
	stream := RepeatFunc(context.Background(), rand)

	require.NotZero(t, <-stream)
	require.NotZero(t, <-stream)

	t.Run("with cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // cancel immediately

		count := 0
		for range RepeatFunc(ctx, func() interface{} { return 0 }) {
			count++
		}

		require.InDelta(t, count, 1, 1)
	})
}
