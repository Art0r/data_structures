package binarysearchtree

import (
	"testing"

	utils "github.com/Art0r/data_structures/utils"
)


func TestBst(t *testing.T) {
	bst := BinarySearchTree{}

	bst.Insert(9)
	bst.Insert(5)
	bst.Insert(1)
	bst.Insert(6)
	bst.Insert(10)
	bst.Insert(11)

	utils.SuccessMessage("ok")
}