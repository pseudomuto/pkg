package collections_test

import (
	"errors"
	"testing"

	. "github.com/pseudomuto/pkg/collections"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	require.Equal(t, []int{2, 4}, Filter(in, func(i int) bool { return i%2 == 0 }))
	require.Equal(t, []int{}, Filter(in, func(i int) bool { return i > len(in) }))

	t.Run("FilterErr", func(t *testing.T) {
		res, err := FilterErr(in, func(i int) (bool, error) { return i%2 == 0, nil })
		require.NoError(t, err)
		require.Equal(t, []int{2, 4}, res)

		res, err = FilterErr(in, func(i int) (bool, error) { return false, errors.New("Boom") })
		require.Nil(t, res)
		require.EqualError(t, err, "Boom")
	})
}

func TestReject(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	require.Equal(t, []int{1, 3, 5}, Reject(in, func(i int) bool { return i%2 == 0 }))
	require.Equal(t, []int{1, 2, 3, 4, 5}, Reject(in, func(i int) bool { return i > len(in) }))

	t.Run("RejectErr", func(t *testing.T) {
		res, err := RejectErr(in, func(i int) (bool, error) { return i%2 == 0, nil })
		require.NoError(t, err)
		require.Equal(t, []int{1, 3, 5}, res)

		res, err = RejectErr(in, func(i int) (bool, error) { return false, errors.New("Boom") })
		require.Nil(t, res)
		require.EqualError(t, err, "Boom")
	})
}
