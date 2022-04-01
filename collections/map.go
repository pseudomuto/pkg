package collections

// Map maps the given slice from type T to type U and returns a new slice.
func Map[T, U any](sl []T, fn func(T) U) []U {
	res, _ := MapErr(sl, func(t T) (U, error) { return fn(t), nil })
	return res
}

// MapErr maps the given slice from type T to type U and returns a new slice. If
// fn returns an error, it will be returned immediately and the mapping operation
// will be stopped.
func MapErr[T, U any](sl []T, fn func(T) (U, error)) ([]U, error) {
	newSl := make([]U, len(sl))
	for i, v := range sl {
		res, err := fn(v)
		if err != nil {
			return nil, err
		}

		newSl[i] = res
	}

	return newSl, nil
}

// MapEntries creates a new map[K2]V2 from the supplied map[K]V by using the
// result of fn for each key/value in m.
func MapEntries[M ~map[K]V, K, K2 comparable, V, V2 any](m M, fn func(K, V) (K2, V2)) map[K2]V2 {
	newMap := make(map[K2]V2)
	for k, v := range m {
		newK, newV := fn(k, v)
		newMap[newK] = newV
	}

	return newMap
}

// MapKeys creates a new map[K2]V from the supplied map[K]V by using the
// result of fn for each key in m.
func MapKeys[M ~map[K]V, K, K2 comparable, V any](m M, fn func(K) K2) map[K2]V {
	res, _ := MapKeysErr(m, func(k K) (K2, error) { return fn(k), nil })
	return res
}

// MapKeysErr creates a new map[K2]V from the supplied map[K]V by using the
// result of fn for each key in m. If fn returns an error, it will be returned
// immediately.
func MapKeysErr[M ~map[K]V, K, K2 comparable, V any](m M, fn func(K) (K2, error)) (map[K2]V, error) {
	newMap := make(map[K2]V)
	for k, v := range m {
		res, err := fn(k)
		if err != nil {
			return nil, err
		}

		newMap[res] = v
	}

	return newMap, nil
}

// MapValues creates a new map[K]V2 from the supplied map[K]V by using the
// result of fn for each value in m.
func MapValues[M ~map[K]V, K comparable, V any, V2 any](m M, fn func(V) V2) map[K]V2 {
	res, _ := MapValuesErr(m, func(v V) (V2, error) { return fn(v), nil })
	return res
}

// MapValuesErr creates a new map[K]V2 from the supplied map[K]V by using the
// result of fn for each value in m. If fn returns an error, it is returned
// immediately.
func MapValuesErr[M ~map[K]V, K comparable, V any, V2 any](m M, fn func(V) (V2, error)) (map[K]V2, error) {
	newMap := make(map[K]V2)
	for k, v := range m {
		res, err := fn(v)
		if err != nil {
			return nil, err
		}

		newMap[k] = res
	}

	return newMap, nil
}
