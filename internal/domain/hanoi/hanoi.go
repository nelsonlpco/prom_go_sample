package hanoi

import "fmt"

type Stack struct {
	data []int
}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Push(item int) {
	s.data = append(s.data, item)
}

func (s *Stack) Pop() int {
	size := len(s.data)

	data := s.data[size-1]
	s.data = s.data[:size-1]

	return data
}

func (s *Stack) Data() []int {
	return s.data
}

func Game() {
	numDiscs := 3
	towerA := NewStack()
	towerB := NewStack()
	towerC := NewStack()

	for i := 1; i <= numDiscs; i++ {
		towerA.Push(i)
	}

	fmt.Println(towerA)
	fmt.Println(towerB)
	fmt.Println(towerC)
	fmt.Println("----------- RESOLVE -----------")

	hanoi(towerA, towerC, towerB, numDiscs)

	fmt.Println(towerA)
	fmt.Println(towerB)
	fmt.Println(towerC)
}

func hanoi(begin, end, temp *Stack, n int) {
	if n == 1 {
		end.Push(begin.Pop())
	} else {
		hanoi(begin, temp, end, n-1)
		hanoi(begin, end, temp, 1)
		hanoi(temp, end, begin, n-1)
	}
}
