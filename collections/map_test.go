package collections_test

import (
	"errors"
	"fmt"
	"testing"

	. "github.com/pseudomuto/pkg/collections"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	in := []int{1, 2, 3, 4, 5}
	exp := []string{"1", "2", "3", "4", "5"}
	require.Equal(t, exp, Map(in, func(i int) string { return fmt.Sprintf("%d", i) }))

	in = []int{}
	exp = []string{}
	require.Equal(t, exp, Map(in, func(i int) string { return fmt.Sprintf("%d", i) }))

	in = nil
	exp = []string{}
	require.Equal(t, exp, Map(in, func(i int) string { return fmt.Sprintf("%d", i) }))

	t.Run("MapErr", func(t *testing.T) {
		in := []int{1, 2, 3, 4, 5}
		res, err := MapErr(in, func(i int) (string, error) { return fmt.Sprintf("%d", i), nil })
		require.NoError(t, err)
		require.Equal(t, []string{"1", "2", "3", "4", "5"}, res)

		res, err = MapErr(in, func(i int) (string, error) { return "", errors.New("Boom") })
		require.Nil(t, res)
		require.EqualError(t, err, "Boom")
	})
}

func TestMapEntries(t *testing.T) {
	in := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	exp := map[string]string{
		"new-one":   "1",
		"new-two":   "2",
		"new-three": "3",
	}

	require.Equal(t, exp, MapEntries(in, func(k string, v int) (string, string) {
		return fmt.Sprintf("new-%s", k), fmt.Sprintf("%d", v)
	}))
}

func TestMapKeys(t *testing.T) {
	in := map[int]int{
		1: 1,
		2: 2,
	}

	require.Equal(t, map[int]int{
		2: 1,
		4: 2,
	}, MapKeys(in, func(i int) int { return i * 2 }))

	require.Equal(t, map[string]int{
		"1": 1,
		"2": 2,
	}, MapKeys(in, func(i int) string { return fmt.Sprintf("%d", i) }))

	t.Run("MapKeysErr", func(t *testing.T) {
		res, err := MapKeysErr(in, func(i int) (int, error) { return i * 2, nil })
		require.NoError(t, err)
		require.Equal(t, map[int]int{
			2: 1,
			4: 2,
		}, res)

		res, err = MapKeysErr(in, func(int) (int, error) { return 0, errors.New("Boom") })
		require.Nil(t, res)
		require.EqualError(t, err, "Boom")
	})
}

func TestMapValues(t *testing.T) {
	in := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
	}

	require.Equal(t, map[string]int{
		"one":   2,
		"two":   4,
		"three": 6,
	}, MapValues(in, func(i int) int { return i * 2 }))

	require.Equal(t, map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
	}, MapValues(in, func(i int) string { return fmt.Sprintf("%d", i) }))

	t.Run("MapValuesErr", func(t *testing.T) {
		res, err := MapValuesErr(in, func(i int) (int, error) { return i * 2, nil })
		require.NoError(t, err)
		require.Equal(t, map[string]int{
			"one":   2,
			"two":   4,
			"three": 6,
		}, res)

		res, err = MapValuesErr(in, func(i int) (int, error) { return 0, errors.New("Boom") })
		require.Nil(t, res)
		require.EqualError(t, err, "Boom")
	})
}
