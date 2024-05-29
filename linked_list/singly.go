package linkedlist

import "fmt"

type Node struct {
	data int
	next *Node
}

type LinkedList struct {
	head  *Node
	tail  *Node
	count int32
}

func (list *LinkedList) Append(value int) {
	new_node := &Node{data: value}

	if list.head == nil {
		list.tail = new_node
		list.head = new_node
	} else {
		list.tail.next = new_node
		list.tail = new_node
	}

	list.count++
}

func (list *LinkedList) Prepend(value int) {
	new_node := &Node{data: value}

	if list.tail == nil {
		list.head = new_node
		list.tail = new_node
	} else {
		new_node.next = list.head
		list.head = new_node
	}

	list.count++
}

func (list *LinkedList) Insert(index int, value int) {

	if index >= int(list.count) || index < 0 {
		return
	}

	if index == 0 {
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

	if index >= int(list.count) || index < 0 {
		return
	}

	if index == 0 {
		list.head = list.head.next
		if list.head == nil {
			list.tail = nil
		}
		list.count--
		return
	}

	leader := list.traverseToIndex(index)
	toRemove := leader.next

	if toRemove == list.tail {
		list.tail = leader
	}

	leader.next = toRemove.next

	list.count--
}

func (list *LinkedList) Lookout() []int {
	currentNode := list.head

	finalList := []int{}
	for currentNode != nil {
		finalList = append(finalList, currentNode.data)
		currentNode = currentNode.next
	}

	return finalList
}

func (list *LinkedList) traverseToIndex(index int) *Node {
	currentNode := list.head
	for i := 0; i < index-1; i++ {
		if currentNode == nil {
			break
		}
		currentNode = currentNode.next
	}
	return currentNode
}

func TestLinkedList() {
	linkedList := LinkedList{}

	linkedList.Prepend(1)
	linkedList.Prepend(2)
	linkedList.Prepend(3)
	linkedList.Prepend(4)

	fmt.Println("-------------------------")

	fmt.Printf("%v\n", linkedList.Lookout())
	fmt.Println("HEAD: ", linkedList.head.data)
	fmt.Println("TAIL: ", linkedList.tail.data)
	fmt.Println("TOTAL: ", linkedList.count)
	fmt.Println("-------------------------")

	linkedList.Insert(2, 5)

	fmt.Printf("%v\n", linkedList.Lookout())
	fmt.Println("HEAD: ", linkedList.head.data)
	fmt.Println("TAIL: ", linkedList.tail.data)
	fmt.Println("TOTAL: ", linkedList.count)
	fmt.Println("-------------------------")

	linkedList.Remove(3)

	fmt.Printf("%v\n", linkedList.Lookout())
	fmt.Println("HEAD: ", linkedList.head.data)
	fmt.Println("TAIL: ", linkedList.tail.data)
	fmt.Println("TOTAL: ", linkedList.count)
	fmt.Println("-------------------------")

	linkedList.Remove(3)

	fmt.Printf("%v\n", linkedList.Lookout())

	fmt.Println("HEAD: ", linkedList.head.data)
	fmt.Println("TAIL: ", linkedList.tail.data)
	fmt.Println("TOTAL: ", linkedList.count)
	fmt.Println("-------------------------")

	linkedList.Append(2)
	linkedList.Append(50)

	fmt.Printf("%v\n", linkedList.Lookout())

	fmt.Println("HEAD: ", linkedList.head.data)
	fmt.Println("TAIL: ", linkedList.tail.data)
	fmt.Println("TOTAL: ", linkedList.count)
	fmt.Println("-------------------------")

	linkedList.Prepend(7)
	linkedList.Prepend(12)

	fmt.Printf("%v\n", linkedList.Lookout())

	fmt.Println("HEAD: ", linkedList.head.data)
	fmt.Println("TAIL: ", linkedList.tail.data)
	fmt.Println("TOTAL: ", linkedList.count)
	fmt.Println("-------------------------")

}