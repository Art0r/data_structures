package binarysearchtree

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

type BinarySearchTree struct {
	root *Node
}

func (bst *BinarySearchTree) Insert(value int) {
	new_node := &Node{data: value}

	if bst.root == nil {
		bst.root = new_node
		return
	}

	current_node := bst.root
	for current_node != nil {

		if current_node.data > value {
			if current_node.left == nil {
				current_node.left = new_node
				break
			}
			current_node = current_node.left
		}

		if current_node.data <= value {
			if current_node.right == nil {
				current_node.right = new_node
				break
			}
			current_node = current_node.right
		}

	}
}

type LookupArgs struct {
	currentNode *Node
	layer int
}

func (bst *BinarySearchTree) Lookup(value int, args *LookupArgs) {
	if args == nil {
		args = &LookupArgs{
			currentNode: bst.root,
			layer: 0,
		}
	}
	
	if args.currentNode.data == value {
		fmt.Println("LAYER: ", args.layer)
		fmt.Println("VALUE: ", args.currentNode.data)
		if  args.currentNode.left != nil {
			fmt.Println("LEFT: ", args.currentNode.left.data)
		}
		if  args.currentNode.right != nil {
			fmt.Println("RIGHT: ", args.currentNode.right.data)
		}
		return
	}

	var goTo *Node 

	if args.currentNode.data > value {
		goTo = args.currentNode.left
	}

	if args.currentNode.data < value {
		goTo = args.currentNode.right
	}

	bst.Lookup(value, &LookupArgs{
		currentNode: goTo,
		layer: args.layer + 1,
	})

}
