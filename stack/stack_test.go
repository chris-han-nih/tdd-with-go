package stack

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type StackSuite struct {
	suite.Suite
}

func TestStackSuite(t *testing.T) {
	suite.Run(t, new(StackSuite))
}

func (s *StackSuite) TestEmpty() {
	stack := NewStack()
	s.True(stack.Empty())
}

func TestEmpty(t *testing.T) {
	s := NewStack()
	if !s.Empty() {
		t.Error("Empty test failed")
	}
}

func (s *StackSuite) TestNotEmpty() {
	stack := NewStack()
	stack.Add(1)
	s.False(stack.Empty())
}

func TestNotEmpty(t *testing.T) {
	s := NewStack()
	s.Add(1)
	if s.Empty() {
		t.Error("Empty test failed")
	}
}

func (s *StackSuite) TestSizeZero() {
	stack := NewStack()
	s.Equal(0, stack.Size())
}

func TestSizeZero(t *testing.T) {
	s := NewStack()
	if s.Size() != 0 {
		t.Error("Size test failed")
	}
}

func (s *StackSuite) TestSizeOne() {
	stack := NewStack()
	stack.Add(1)
	s.Equal(1, stack.Size())
}

func TestSizeOne(t *testing.T) {
	s := NewStack()
	s.Add(1)
	if s.Size() != 1 {
		t.Error("Size test failed")
	}
}

func (s *StackSuite) TestSizeTwo() {
	stack := NewStack()
	stack.Add(1)
	stack.Add(2)
	s.Equal(2, stack.Size())
}

func TestSizeThree(t *testing.T) {
	s := NewStack()
	s.Add(1)
	s.Add(2)
	s.Add(3)
	if s.Size() != 3 {
		t.Error("Size test failed")
	}
}

func (s *StackSuite) TestPopOne() {
	stack := NewStack()
	stack.Add(1)

	v := stack.Pop()

	s.Equal(1, v)
}

func TestPopOne(t *testing.T) {
	s := NewStack()
	s.Add(1)

	v := s.Pop()

	if v != 1 {
		t.Error("Expected 1, got ", v)
	}

	if s.Size() != 0 {
		t.Error("Expected size 0, got ", s.Size())
	}
}

func (s *StackSuite) TestPopTwo() {
	stack := NewStack()
	stack.Add(1)
	stack.Add(2)

	v1 := stack.Pop()

	s.Equal(2, v1)

	s.Equal(1, stack.Size())

	v2 := stack.Pop()

	s.Equal(1, v2)

	s.Equal(0, stack.Size())
}

func TestPopTwo(t *testing.T) {
	s := NewStack()
	s.Add(1)
	s.Add(2)

	v1 := s.Pop()

	if v1 != 2 {
		t.Error("Expected 2, got ", v1)
	}

	if s.Size() != 1 {
		t.Error("Expected size 1, got ", s.Size())
	}

	v2 := s.Pop()

	if v2 != 1 {
		t.Error("Expected 1, got ", v2)
	}

	if s.Size() != 0 {
		t.Error("Expected size 0, got ", s.Size())
	}
}
