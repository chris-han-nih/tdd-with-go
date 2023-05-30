package stack

import "testing"

func TestEmpty(t *testing.T) {
	var s Stack
	if !s.Empty() {
		t.Error("Empty test failed")
	}
}
