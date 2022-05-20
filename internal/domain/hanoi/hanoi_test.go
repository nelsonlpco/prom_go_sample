package hanoi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Stack(t *testing.T) {
	stack := NewStack()

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	v := stack.Pop()
	assert.Equal(t, 3, v)

	v = stack.Pop()
	assert.Equal(t, 2, v)

	v = stack.Pop()
	assert.Equal(t, 1, v)
}

func Test_Hanoi(t *testing.T) {
	numDiscs := 3
	towerA := NewStack()
	towerB := NewStack()
	towerC := NewStack()

	for i := 1; i <= numDiscs; i++ {
		towerA.Push(i)
	}

	hanoi(towerA, towerC, towerB, numDiscs)

	assert.Equal(t, []int{}, towerA.data)
	assert.Equal(t, []int{}, towerB.data)
	assert.Equal(t, []int{1, 2, 3}, towerC.data)
}
