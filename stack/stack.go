package stack

type Stack struct {
	size   int
	values []int
}

func NewStack() *Stack {
	return &Stack{0, []int{}}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Add(i int) {
	s.size++
	s.values = append(s.values, i)
}

func (s *Stack) Pop() int {
	s.size--
	val := s.values[s.size]
	s.values = s.values[:s.size]
	return val
}
