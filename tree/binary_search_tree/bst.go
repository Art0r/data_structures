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

	newNode := &Node{data: value}

	if bst.root == nil {
		bst.root = newNode
		return
	}

	currentNode := bst.root
	for currentNode != nil {
		if currentNode.data > value {
			if currentNode.left == nil {
				currentNode.left = newNode
				break
			}
			currentNode = currentNode.left
		}

		if currentNode.data <= value {
			if currentNode.right == nil {
				currentNode.right = newNode
				break
			}
			currentNode = currentNode.right
		}
	}

}

func (bst *BinarySearchTree) Lookup(value int) *Node {

	node, _ := bst.traverse(value, bst.root, nil)

	return node
}

func (bst *BinarySearchTree) Remove(value int) {

	currentNode, parentNode := bst.traverse(value, bst.root, nil)

	if currentNode == nil || parentNode == nil {
		return
	}

	// leaf node (no child)
	if currentNode.left == nil && currentNode.right == nil {
		if parentNode.left == currentNode {
			currentNode = nil
			parentNode.left = nil
		}

		if parentNode.right == currentNode {
			currentNode = nil
			parentNode.right = nil
		}
		return
	}

	// parent with one child
	if currentNode.left == nil || currentNode.right == nil {
		if parentNode.left == currentNode {
			parentNode.left = currentNode.left
			currentNode.left = nil
		}

		if parentNode.right == currentNode {
			parentNode.right = currentNode.right
			currentNode.right = nil
		}
		return
	}

	// parent with two childs
	// TODO stills cases to cover for this condition
	nodeToReplace := currentNode.right

	for {
		if nodeToReplace.left == nil {
			break
		}
		nodeToReplace = nodeToReplace.left
	}

	if parentNode.left == currentNode {
		parentNode.left = nodeToReplace
		parentNode.left.left = currentNode.left
		currentNode.left = nil
	}

	if parentNode.right == currentNode {
		parentNode.right = nodeToReplace
		parentNode.right.right = currentNode.right
		currentNode.right = nil
	}
}

func (bst *BinarySearchTree) traverse(value int, currentNode *Node, parentNode *Node) (*Node, *Node) {

	if currentNode == nil {
		return nil, nil
	}

	if currentNode.data == value {
		return currentNode, parentNode
	}

	if currentNode.data > value {
		return bst.traverse(value, currentNode.left, currentNode)
	}

	if currentNode.data < value {
		return bst.traverse(value, currentNode.right, currentNode)
	}

	return nil, nil
}

func (bst *BinarySearchTree) printInOrder(node *Node) {
	if node != nil {
		bst.printInOrder(node.left)
		fmt.Printf("%d ", node.data)
		bst.printInOrder(node.right)
	}
}

// PrintTree prints the entire BST.
func (bst *BinarySearchTree) PrintTree() {
	bst.printInOrder(bst.root)
	fmt.Println()
}
