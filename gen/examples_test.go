package gen_test

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/pseudomuto/pkg/gen"
)

// Print the first 5 integers
func ExampleRange() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := range gen.Range(ctx, 5) {
		fmt.Println(i)
	}

	// Output:
	// 0
	// 1
	// 2
	// 3
	// 4
}

// Read values from the generator. In this example we supply 3 values to repeat and read from the channel 10 times. This
// shows that successive reads from the channel wrap the input values.
func ExampleRepeat() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream := gen.Repeat(ctx, 1, 2, 3)
	for i := 0; i < 10; i++ {
		fmt.Println(<-stream)
	}

	// Output:
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
	// 1
	// 2
	// 3
	// 1
}

// Print 5 random integers
func ExampleRepeatFunc() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rand := func() int { return rand.Int() }
	stream := gen.RepeatFunc(ctx, rand)

	for i := 0; i < 5; i++ {
		fmt.Println(<-stream)
	}
}

// Print the number 1 ten times
func ExampleTake() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	for i := range gen.Take(ctx, gen.Repeat(ctx, 1), 10) {
		fmt.Print(i)
	}

	// Output: 1111111111
}
