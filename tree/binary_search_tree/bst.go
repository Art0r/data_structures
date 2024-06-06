package binarysearchtree

type Node struct {
	data int
	left *Node
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

func (bst *BinarySearchTree) Lookup(value int) {
	// verificar se um valor existe na arvore
	current_node := bst.root

	for current_node != nil {

	}
}