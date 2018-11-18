package gen_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	. "github.com/pseudomuto/pkg/gen"
	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {
	tests := []struct {
		inputs  []interface{}
		outputs []interface{}
	}{
		{
			inputs:  []interface{}{1},
			outputs: []interface{}{1, 1, 1, 1, 1},
		},
		{
			inputs:  []interface{}{1, 2, 3},
			outputs: []interface{}{1, 2, 3, 1, 2, 3, 1, 2, 3, 1},
		},
		{
			inputs:  []interface{}{"a"},
			outputs: []interface{}{"a", "a", "a"},
		},
	}

	ctx := context.Background()
	for _, test := range tests {
		stream := Repeat(ctx, test.inputs...)

		for _, val := range test.outputs {
			assert.Equal(t, val, <-stream)
		}
	}

	t.Run("with cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // cancel immediately

		count := 0
		for range Repeat(ctx, 1) {
			count++
		}

		assert.InDelta(t, count, 1, 1)
	})
}

// Read values from the generator. In this example we supply 3 values to repeat and read from the channel 10 times. This
// shows that successive reads from the channel wrap the input values.
func ExampleRepeat() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	stream := Repeat(ctx, 1, 2, 3)
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

func TestRepeatFunc(t *testing.T) {
	rand := func() interface{} { return rand.Int() }
	stream := RepeatFunc(context.Background(), rand)

	assert.NotZero(t, <-stream)
	assert.NotZero(t, <-stream)

	t.Run("with cancellation", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())
		cancel() // cancel immediately

		count := 0
		for range RepeatFunc(ctx, func() interface{} { return 0 }) {
			count++
		}

		assert.InDelta(t, count, 1, 1)
	})
}

// Print 5 random integers
func ExampleRepeatFunc() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	rand := func() interface{} { return rand.Int() }
	stream := RepeatFunc(ctx, rand)

	for i := 0; i < 5; i++ {
		fmt.Println(<-stream)
	}
}
