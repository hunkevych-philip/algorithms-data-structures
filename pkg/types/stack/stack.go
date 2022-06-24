package stack

import "fmt"

type Stack struct {
	Top    *Node
	Bottom *Node
	Length int
}

func (s *Stack) Peek() interface{} {
	if s.Top == nil {
		fmt.Println("stack ended")
		return 0
	}

	return s.Top.Val
}

// Push saves new node on a top of a stack
func (s *Stack) Push(v interface{}) {
	defer func() {
		s.Length++
	}()

	newNode := &Node{
		Val: v,
	}

	if s.Top == nil {
		s.Top = newNode
		s.Bottom = newNode
	} else {
		holding := s.Top
		s.Top = newNode
		s.Top.Next = holding
	}
}

// Pop removes the top node
func (s *Stack) Pop() {
	if s.Length == 0 {
		return
	}

	defer func() {
		s.Length--
	}()

	s.Top = s.Top.Next
	if s.Top == nil {
		s.Bottom = nil
	}
}
