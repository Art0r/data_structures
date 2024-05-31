package utils

import (
	"fmt"
	"testing"
)


func ProccessFunc(msg string, err error, t *testing.T) {
	if err != nil {
		t.Error(err)
	}
	t.Log(msg)
}

func ErrorMessage(name string, expected []int, received []int, count int) error {
	return fmt.Errorf("[%s] failed, expected: %v but received %v. Expected count: %d", name, expected, received, count)
}

func SuccessMessage(name string) string {
	return fmt.Sprintf("[%s] succeeded", name)
}