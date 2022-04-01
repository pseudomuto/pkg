package collections_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	. "github.com/pseudomuto/pkg/collections"
	"github.com/stretchr/testify/require"
)

func TestAll(t *testing.T) {
	in := []string{"one", "two", "three"}
	require.False(t, All(in, func(s string) bool { return strings.HasPrefix(s, "t") }))
	require.True(t, All(in, func(s string) bool { return len(s) > 1 }))

	t.Run("AllErr", func(t *testing.T) {
		res, err := AllErr(in, func(s string) (bool, error) { return len(s) > 1, nil })
		require.NoError(t, err)
		require.True(t, res)

		res, err = AllErr(in, func(s string) (bool, error) { return true, errors.New("Boom") })
		require.False(t, res)
		require.EqualError(t, err, "Boom")
	})
}

func TestFold(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	require.Equal(t, 15, Fold(in, 0, func(i, acc int) int { return acc + i }))

	exp := map[string]int{
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
	}

	require.Equal(t, exp, Fold(in, make(map[string]int), func(i int, acc map[string]int) map[string]int {
		acc[fmt.Sprintf("%d", i)] = i
		return acc
	}))

	t.Run("FoldErr", func(t *testing.T) {
		res, err := FoldErr(in, 0, func(i, acc int) (int, error) { return acc + i, nil })
		require.NoError(t, err)
		require.Equal(t, 15, res)

		res, err = FoldErr(in, 0, func(i, acc int) (int, error) { return acc + 1, errors.New("Boom") })
		require.Equal(t, 0, res)
		require.EqualError(t, err, "Boom")
	})
}

func TestNone(t *testing.T) {
	in := []string{"one", "two", "three"}
	require.True(t, None(in, func(s string) bool { return strings.HasPrefix(s, "0") }))
	require.False(t, None(in, func(s string) bool { return len(s) > 1 }))

	t.Run("NoneErr", func(t *testing.T) {
		res, err := NoneErr(in, func(s string) (bool, error) { return len(s) == 1, nil })
		require.NoError(t, err)
		require.True(t, res)

		res, err = NoneErr(in, func(s string) (bool, error) { return true, errors.New("Boom") })
		require.False(t, res)
		require.EqualError(t, err, "Boom")
	})
}

func TestSum(t *testing.T) {
	ints := []int{1, 2, 3, 4, 5}
	require.Equal(t, 15, Sum(ints))

	u8 := []uint8{}
	require.Equal(t, uint8(0), Sum(u8))

	floats := []float64{1, 2, 3, 4, 5}
	require.Equal(t, 15.0, Sum(floats))
}
