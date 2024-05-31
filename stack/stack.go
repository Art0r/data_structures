package stack


type Stack struct {
	Top    *Node
	Bottom *Node
	Count  int32
}

func (stack *Stack) Push(value int) {
	new_node := &Node{data: value}

	if stack.Top == nil {
		stack.Top = new_node
		stack.Bottom = new_node
		stack.Count++
		return
	}

	new_node.next = stack.Top
	stack.Top = new_node
	stack.Count++
}

func (stack *Stack) Pop() {

	if stack.Top == nil || stack.Bottom == nil {
		return
	}

	stack.Top = stack.Top.next

	stack.Count--
}

func (stack *Stack) Peek(index int) int {
	if stack.Count == 0 {
		return -1
	}

	node := stack.transverseStack(index)
	return node.data
}

func (stack *Stack) Lookout() []int {
	currentNode := stack.Top

	finalList := []int{}
	for currentNode != nil {
		finalList = append(finalList, currentNode.data)
		currentNode = currentNode.next
	}

	return finalList
}

func (stack *Stack) transverseStack(index int) *Node {
	currentNode := stack.Top

	for i := 0; i < index-1; i++ {
		if currentNode == nil {
			break
		}
		currentNode = currentNode.next
	}

	return currentNode
}