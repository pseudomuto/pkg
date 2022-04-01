package collections_test

import (
	"fmt"
	"strconv"

	"github.com/pseudomuto/pkg/collections"
)

// Count occurrences of works in a slice
func ExampleFold() {
	in := []string{"some", "words", "go", "here", "and", "here", "words"}
	counts := collections.Fold(in, make(map[string]int), func(word string, acc map[string]int) map[string]int {
		if _, ok := acc[word]; !ok {
			acc[word] = 0
		}

		acc[word]++
		return acc
	})

	for _, w := range []string{"and", "go", "here", "some", "words"} {
		fmt.Printf("%s: %d\n", w, counts[w])
	}

	// Output:
	// and: 1
	// go: 1
	// here: 2
	// some: 1
	// words: 2
}

// Convert slice of int to slice of string
func ExampleMap() {
	strings := collections.Map([]int{1, 2, 3, 4, 5}, func(i int) string {
		return fmt.Sprintf("str-%d", i)
	})

	for _, str := range strings {
		fmt.Println(str)
	}

	// Output:
	// str-1
	// str-2
	// str-3
	// str-4
	// str-5
}

// Convert slice of string to slice of int using strconv.
func ExampleMapErr() {
	// all good, no errors here
	ints, _ := collections.MapErr([]string{"10", "20", "30"}, strconv.Atoi)

	for _, i := range ints {
		fmt.Printf("%d\n", i+1)
	}

	// can't parse "nope" into int
	_, err := collections.MapErr([]string{"nope"}, strconv.Atoi)
	fmt.Println(err)

	// Output:
	// 11
	// 21
	// 31
	// strconv.Atoi: parsing "nope": invalid syntax
}

// Sum of up the the first 5 integers
func ExampleSum() {
	fmt.Printf("%d\n", collections.Sum([]int{1, 2, 3, 4, 5}))
	// Output: 15
}
