package stack

import "testing"

func TestEmpty(t *testing.T) {
	s := NewStack()
	if !s.Empty() {
		t.Error("Empty test failed")
	}
}

func TestNotEmpty(t *testing.T) {
	s := NewStack()
	s.Add(1)
	if s.Empty() {
		t.Error("Empty test failed")
	}
}

func TestSizeZero(t *testing.T) {
	s := NewStack()
	if s.Size() != 0 {
		t.Error("Size test failed")
	}
}

func TestSizeOne(t *testing.T) {
	s := NewStack()
	s.Add(1)
	if s.Size() != 1 {
		t.Error("Size test failed")
	}
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
