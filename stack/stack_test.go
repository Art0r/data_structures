package stack

import (
	"reflect"
	"testing"

	utils "github.com/Art0r/data_structures/utils"
)


func TestStack(t *testing.T) {

	s := Stack{}

	if s.Top != nil || s.Bottom != nil || s.Count != 0 {
		t.Error("Stack is not initially empty")
	}

	msg, err := testPush(&s)
	utils.ProccessFunc(msg, err, t)

	msg, err = testPop(&s)
	utils.ProccessFunc(msg, err, t)

	msg, err = testPopTilEmpty(&s)
	utils.ProccessFunc(msg, err, t)

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	msg, err = testPeek(&s)
	utils.ProccessFunc(msg, err, t)

	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()

	msg, err = testPeekEmpty(&s)
	utils.ProccessFunc(msg, err, t)

	msg, err = testPopEmpty(&s)
	utils.ProccessFunc(msg, err, t)

}

func testPopEmpty(s *Stack) (string, error) {
	s.Pop()
	return utils.SuccessMessage("testPopEmpty"), nil
}


func testPeekEmpty(s *Stack) (string, error) {
	result := s.Peek(1)
	expected := -1

	if result != expected {
		return "", utils.ErrorMessage("testPeek", []int{expected}, []int{result}, int(s.Count))
	}
	return utils.SuccessMessage("testPeekEmpty"), nil
}

func testPeek(s *Stack) (string, error) {
	result := s.Peek(1)
	expected := 4

	if result != expected {
		return "", utils.ErrorMessage("testPeek", []int{expected}, []int{result}, int(s.Count))
	}
	return utils.SuccessMessage("testPeek"), nil
}

func testPopTilEmpty(s *Stack) (string, error) {
	s.Pop()
	s.Pop()
	s.Pop()
	s.Pop()

	lookout := s.Lookout()
	expected := []int{}

	if !reflect.DeepEqual(lookout, expected) && s.Count == 0 {
		return "", utils.ErrorMessage("testPopTilEmpty", expected, lookout, int(s.Count))
	}
	return utils.SuccessMessage("testPopTilEmpty"), nil
}

func testPop(s *Stack) (string, error) {
	s.Pop()
	s.Pop()

	lookout := s.Lookout()
	expected := []int{2, 1}

	if !reflect.DeepEqual(lookout, expected) && s.Count == 2 {
		return "", utils.ErrorMessage("testPop", expected, lookout, int(s.Count))
	}
	return utils.SuccessMessage("testPop"), nil
}

func testPush(s *Stack) (string, error) {
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	lookout := s.Lookout()
	expected := []int{4, 3, 2, 1}

	if !reflect.DeepEqual(lookout, expected) && s.Count == 4 {
		return "", utils.ErrorMessage("testPush", expected, lookout, int(s.Count))
	}
	return utils.SuccessMessage("testPush"), nil
}
