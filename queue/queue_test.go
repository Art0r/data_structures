package queue

import (
	"reflect"
	"testing"

	utils "github.com/Art0r/data_structures/utils"
)


func TestQueue(t *testing.T) {

	q := Queue{}

	if q.Head != nil || q.Tail != nil || q.Count != 0 {
		t.Error("Queue is not initially empty")
	}

	msg, err := testPush(&q)
	utils.ProccessFunc(msg, err, t)

	msg, err = testPop(&q)
	utils.ProccessFunc(msg, err, t)

	msg, err = testPopTilEmpty(&q)
	utils.ProccessFunc(msg, err, t)

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	msg, err = testPeek(&q)
	utils.ProccessFunc(msg, err, t)

	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	msg, err = testPeekEmpty(&q)
	utils.ProccessFunc(msg, err, t)

	msg, err = testPopEmpty(&q)
	utils.ProccessFunc(msg, err, t)

}

func testPopEmpty(q *Queue) (string, error) {
	q.Dequeue()
	return utils.SuccessMessage("testDequeueEmpty"), nil
}


func testPeekEmpty(q *Queue) (string, error) {
	result := q.Peek(1)
	expected := -1

	if result != expected {
		return "", utils.ErrorMessage("testPeek", []int{expected}, []int{result}, int(q.Count))
	}
	return utils.SuccessMessage("testPeekEmpty"), nil
}

func testPeek(q *Queue) (string, error) {
	result := q.Peek(1)
	expected := 4

	if result != expected {
		return "", utils.ErrorMessage("testPeek", []int{expected}, []int{result}, int(q.Count))
	}
	return utils.SuccessMessage("testPeek"), nil
}

func testPopTilEmpty(q *Queue) (string, error) {
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()
	q.Dequeue()

	lookout := q.Lookout()
	expected := []int{}

	if !reflect.DeepEqual(lookout, expected) && q.Count == 0 {
		return "", utils.ErrorMessage("testDequeueTilEmpty", expected, lookout, int(q.Count))
	}
	return utils.SuccessMessage("testDequeueTilEmpty"), nil
}

func testPop(q *Queue) (string, error) {
	q.Dequeue()
	q.Dequeue()

	lookout := q.Lookout()
	expected := []int{2, 1}

	if !reflect.DeepEqual(lookout, expected) && q.Count == 2 {
		return "", utils.ErrorMessage("testDequeue", expected, lookout, int(q.Count))
	}
	return utils.SuccessMessage("testDequeue"), nil
}

func testPush(q *Queue) (string, error) {
	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)
	q.Enqueue(4)

	lookout := q.Lookout()
	expected := []int{4, 3, 2, 1}

	if !reflect.DeepEqual(lookout, expected) && q.Count == 4 {
		return "", utils.ErrorMessage("testEnqueue", expected, lookout, int(q.Count))
	}
	return utils.SuccessMessage("testEnqueue"), nil
}
