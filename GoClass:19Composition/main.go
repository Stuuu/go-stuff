package main

import (
	"fmt"
)

type StringStack struct {
	data []string
}

func (s *StringStack) Push(x string) {
	s.data = append(s.data, x)
}

func (s *StringStack) Pop() string {
	if l := len(s.data); l > 0 {
		t := s.data[l-1]
		s.data = s.data[:l-1]
		return t
	}
	
	panic("pop from empty stack")
}
func main() {

	stack := StringStack{}
	
	stack.Push("test1")
	stack.Push("test2")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
}
