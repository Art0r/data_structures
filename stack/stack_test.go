package stack

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	
    s := Stack{}

	if s.Top != nil || s.Bottom != nil || s.Count != 0 {
		t.Error("Stack is not initially empty")
	}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	first_lookout := s.Lookout()
	expected := []int{4, 3, 2, 1}

	if  !reflect.DeepEqual(first_lookout, expected) {
		t.Errorf("Push failed, expected: %v but received %v", expected, first_lookout)
	}
}

