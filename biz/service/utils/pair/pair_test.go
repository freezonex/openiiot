package pair

import "testing"

func TestNewPrimitives(t *testing.T) {
	p := New(1, "one")
	if p.First != 1 || p.Second != "one" {
		t.Error()
	}
}

func TestNewPointers(t *testing.T) {
	left, right := 1, "one"
	p := New(&left, &right)
	if p.First != &left || p.Second != &right {
		t.Error()
	}
}
