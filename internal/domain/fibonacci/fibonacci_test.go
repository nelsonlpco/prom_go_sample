package fibonacci

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fibonacci(t *testing.T) {
	scenes := []struct {
		i uint64
		o uint64
	}{
		{i: 0, o: 0},
		{i: 1, o: 1},
		{i: 2, o: 1},
		{i: 3, o: 2},
		{i: 4, o: 3},
		{i: 5, o: 5},
		{i: 6, o: 8},
		{i: 7, o: 13},
		{i: 8, o: 21},
	}

	fib := New()

	for _, scene := range scenes {
		t.Run(fmt.Sprintf("when receive %v returns: %v", scene.i, scene.o), func(t *testing.T) {
			result := fib.Calc(scene.i)

			assert.Equal(t, scene.o, result)
		})
	}
}
