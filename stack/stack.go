package stack

type Stack struct{}

func NewStack() *Stack {
	return new(Stack)
}

func (s *Stack) Empty() bool {
	return true
}
