package binarysearchtree

import (
	"fmt"
	"testing"

)


func TestBst(t *testing.T) {
	bst := BinarySearchTree{}

	bst.Insert(9)
	bst.Insert(5)
	bst.Insert(1)
	bst.Insert(6)
	bst.Insert(10)
	bst.Insert(11)

	/*
	     9
		 /\
		/  \
	   5   10
	  / \    \
	 /   \	  \
	1     6   11


	*/

	node := bst.Lookup(5)
	if node != nil {
		fmt.Println("VALUE: ", node.data)
		if node.left != nil {
			fmt.Println("LEFT: ", node.left.data)	
		}
		if node.right != nil {
			fmt.Println("RIGHT: ", node.right.data)	
		}
	} else {
		fmt.Println("VALUE NOT FOUND")
	}

	fmt.Println("-------------------------------------")
	bst.PrintTree()

	bst.Remove(5)

	fmt.Println("-------------------------------------")
	bst.PrintTree()

	
	// fmt.Println("-------------------------------------")
	// bst.PrintTree()

	// bst.Remove(5)
	
	// fmt.Println("-------------------------------------")
	// bst.PrintTree()
}