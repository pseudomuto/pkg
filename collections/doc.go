// Package collections provides utilities for working with collections (slices
// and maps).
//
// Things like Map, Fold, Filter, Reject, etc. are available as generic top-level
// functions. In almost every case, there is an `<func>Err` equivalent (MapErr,
// FoldErr, etc.), which allows for transform/predicate functions to return an error.
//
// Finally, all of these methods return new objects rather than mutate that target.
package collections
