package collections

import "golang.org/x/exp/constraints"

// All determines whether or not the predicate returns true for all elements in
// sl.
func All[T any](sl []T, pred func(T) bool) bool {
	res, _ := AllErr(sl, func(t T) (bool, error) { return pred(t), nil })
	return res
}

// AllErr determines whether or not the predicate is true for all members in sl.
// If pred returns an error, it is returned immediately and the condition is
// deemed to be false.
func AllErr[T any](sl []T, pred func(T) (bool, error)) (bool, error) {
	for _, v := range sl {
		ok, err := pred(v)
		if err != nil {
			return false, err
		}

		if !ok {
			return false, nil
		}
	}

	return true, nil
}

// Fold creates a single value of type U from the slice of T. It works by
// iterating over all elements in sl calling fn for each value along with
// the accumlator (starts with initVal).
//
// Each call to fn returns the updated accumulator. At the end of the
// iteration, the final value is returned.
func Fold[T, U any](sl []T, initVal U, fn func(T, U) U) U {
	res, _ := FoldErr(sl, initVal, func(t T, u U) (U, error) { return fn(t, u), nil })
	return res
}

// FoldErr creates a single value of type U from the slice of T. It works by
// iterating over all elements in sl calling fn for each value along with
// the accumlator (starts with initVal). If fn returns an error, it will return
// the initVal value and the error immediately without continuing the fold.
//
// Each call to fn returns the updated accumulator. At the end of the
// iteration, the final value is returned.
func FoldErr[T, U any](sl []T, initVal U, fn func(T, U) (U, error)) (U, error) {
	var err error
	val := initVal
	for _, t := range sl {
		val, err = fn(t, val)
		if err != nil {
			return initVal, err
		}
	}

	return val, nil
}

// None determines whether the predicate is false for all members of sl.
func None[T any](sl []T, pred func(T) bool) bool {
	res, _ := NoneErr(sl, func(t T) (bool, error) { return pred(t), nil })
	return res
}

// NoneErr determines whether the predicate is false for all members of sl. If
// pred returns an error, the error is returned immediately and the result is
// deemed to be false.
func NoneErr[T any](sl []T, pred func(T) (bool, error)) (bool, error) {
	return AllErr(sl, func(t T) (bool, error) {
		ok, err := pred(t)
		return !ok, err
	})
}

// Sum sums up all of the values in sl and returns the result.
func Sum[T constraints.Integer | constraints.Float](sl []T) T {
	return Fold(sl, 0, func(t T, acc T) T { return acc + t })
}
