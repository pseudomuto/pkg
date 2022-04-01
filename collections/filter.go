package collections

// Filter returns a new slice of T containing only those elements for which pred
// returns true.
func Filter[T any](sl []T, pred func(T) bool) []T {
	res, _ := FilterErr(sl, func(t T) (bool, error) { return pred(t), nil })
	return res
}

// FilterErr returns a new slice of T containing only those elements for which
// pred returns true. If pred returns an error, it will return that error
// immediately and discontinue iteration.
func FilterErr[T any](sl []T, pred func(T) (bool, error)) ([]T, error) {
	res := make([]T, 0, len(sl))
	for _, v := range sl {
		ok, err := pred(v)
		if err != nil {
			return nil, err
		}

		if ok {
			res = append(res, v)
		}
	}

	// return the "clipped" slice
	return res[:len(res):len(res)], nil
}

// Reject returns a new slice of T containing only those element for which pred
// returns false.
func Reject[T any](sl []T, pred func(T) bool) []T {
	res, _ := RejectErr(sl, func(t T) (bool, error) { return pred(t), nil })
	return res
}

// RejectErr returns a new slice of T containing only those element for which
// pred returns false. If pred returns an error, it will be returned immediately
// and iteration will be stopped.
func RejectErr[T any](sl []T, pred func(T) (bool, error)) ([]T, error) {
	return FilterErr(sl, func(t T) (bool, error) {
		ok, err := pred(t)
		return !ok, err
	})
}
