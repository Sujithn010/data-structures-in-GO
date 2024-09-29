package main

import "fmt"

type Stack struct {
	top    int
	values []int
}

func NewStack(capacity int) *Stack {
	return &Stack{
		top:    -1,
		values: make([]int, capacity),
	}
}

func (s Stack) isEmpty() bool {
	return s.top == -1
}

func (s *Stack) push(val int) {
	s.top++
	s.values[s.top] = val
}

func (s *Stack) pop() int {
	if s.isEmpty() {
		fmt.Println("Stack in empty")
		return -1
	}

	retVal := s.values[s.top]
	s.top--
	return retVal
}

func (s Stack) printStack() {
	fmt.Println("Stack elements:")
	for i := s.top; i >= 0; i-- {
		fmt.Println(s.values[i])
	}
}

func main() {
	stk := NewStack(2)
	for i := 0; i < 2; i++ {
		fmt.Printf("Pushing element:%d\n", i+1)
		stk.push(i + 1)
	}

	// stack is not empty here
	fmt.Printf("Stack TOP:%d\n", stk.values[stk.top])

	fmt.Println("Printing stack elements:-")
	stk.printStack()

	fmt.Println("Popping element from stack-")
	fmt.Printf("Popped element:%d\n", stk.pop())

	fmt.Println("Popping element from stack-")
	fmt.Printf("Popped element:%d\n", stk.pop())

}
