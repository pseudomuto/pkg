package gen_test

import (
	"context"
	"testing"

	. "github.com/pseudomuto/pkg/gen"
	"github.com/stretchr/testify/require"
)

func TestRange(t *testing.T) {
	tests := []struct {
		n        int
		expected []int
	}{
		{n: 10, expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{n: 0, expected: []int{}},
		{n: -1, expected: []int{}},
	}

	ctx := context.Background()
	for _, test := range tests {
		results := make([]int, 0, len(test.expected))
		for val := range Range(ctx, test.n) {
			results = append(results, val)
		}

		require.Equal(t, test.expected, results)
	}

	t.Run("with cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // cancel immediately

		count := 0
		for range Range(ctx, 10) {
			count++
		}

		require.InDelta(t, count, 1, 1)
	})
}

func TestRangeFrom(t *testing.T) {
	tests := []struct {
		start    int
		n        int
		expected []int
	}{
		{n: 10, expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{start: 1, n: 10, expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{start: 1000, n: 3, expected: []int{1000, 1001, 1002}},
		{n: 0, expected: []int{}},
		{n: -1, expected: []int{}},
	}

	ctx := context.Background()
	for _, test := range tests {
		results := make([]int, 0, len(test.expected))
		for val := range RangeFrom(ctx, test.start, test.n) {
			results = append(results, val)
		}

		require.Equal(t, test.expected, results)
	}

	t.Run("with cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // cancel immediately

		count := 0
		for range RangeFrom(ctx, 1, 10) {
			count++
		}

		require.InDelta(t, count, 1, 1)
	})
}
