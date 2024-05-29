package main

import "fmt"

type Node struct {
	data any
	next *Node
}

type LinkedList struct {
	head  *Node
	tail  *Node
	count int32
}

func (list *LinkedList) Append(value any) {
	new_node := &Node{data: value}

	if list.head == nil {
		list.head = new_node
	} else {
		previous_head := list.head
		list.head = new_node
		list.head.next = previous_head
	}

	list.count++
}


func (list *LinkedList) Insert(index int, value any) {

	if index >= int(list.count) {
		list.Append(value)
		return
	}

	newNode := &Node{data: value}
	
	leader := list.traverseToIndex(index - 1)
	holder := leader.next

	leader.next = newNode
	leader.next.next = holder

	list.count++
}

func (list *LinkedList) Remove(index int) {
	
	if index >= int(list.count) {
		return
	}

	leader := list.traverseToIndex(index - 1)
	toRemove := leader.next

	leader.next = toRemove.next

	list.count--
}

func (list *LinkedList) Lookout() []any {
	currentNode := list.head

	finalList := []any{}
	for currentNode != nil {
		finalList = append(finalList, currentNode.data)
		currentNode = currentNode.next
	}

	return finalList
}

func (list *LinkedList) traverseToIndex(index int) *Node {
	currentNode := list.head
	for i := 0; i < index; i++ {
		if currentNode == nil {
			break
		}
		currentNode = currentNode.next
	}
	return currentNode
}

func main() {
	linkedList := LinkedList{}

	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)

	fmt.Printf("%v\n", linkedList.Lookout())

	linkedList.Insert(2, 5)

	fmt.Printf("%v\n", linkedList.Lookout())

	linkedList.Remove(3)

	fmt.Printf("%v\n", linkedList.Lookout())

}
