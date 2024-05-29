package stack

import "fmt"

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

func Test() {
	stack := Stack{}

	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)

	fmt.Printf("%v\n", stack.Lookout())

	stack.Pop()

	fmt.Printf("%v\n", stack.Lookout())

	fmt.Printf("%v\n", stack.Peek(3))
}
