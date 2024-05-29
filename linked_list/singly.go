package linkedlist

import "fmt"

type SLinkedList struct {
	head  *Node
	tail  *Node
	count int32
}

func (list *SLinkedList) Append(value int) {
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

func (list *SLinkedList) Prepend(value int) {
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

func (list *SLinkedList) Insert(index int, value int) {

	if index >= int(list.count) || index < 0 {
		return
	}

	if index == 0 {
		list.Append(value)
		return
	}

	newNode := &Node{data: value}

	leader := list.traverseToIndex(index)
	holder := leader.next

	leader.next = newNode
	leader.next.next = holder

	list.count++
}

func (list *SLinkedList) Remove(index int) {

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

func (list *SLinkedList) Lookout() []int {
	currentNode := list.head

	finalList := []int{}
	for currentNode != nil {
		finalList = append(finalList, currentNode.data)
		currentNode = currentNode.next
	}

	return finalList
}

func (list *SLinkedList) Reverse() {
	if list.count == 0 || list.count == 1 {
		return
	}

	first := list.head
	second := first.next

	for second != nil {
		tmp := second.next

		second.next = first
		first = second
		second = tmp
	}

	list.head.next = nil
	list.head = first
}

func (list *SLinkedList) traverseToIndex(index int) *Node {
	currentNode := list.head
	for i := 0; i < index-1; i++ {
		if currentNode == nil {
			break
		}
		currentNode = currentNode.next
	}
	return currentNode
}

func TestSLinkedListReverse() {
	linkedList := SLinkedList{}

	linkedList.Prepend(1)
	linkedList.Prepend(2)
	linkedList.Prepend(3)
	linkedList.Prepend(4)
	linkedList.Prepend(5)
	linkedList.Prepend(6)

	fmt.Println("-------------------------")

	fmt.Printf("%v\n", linkedList.Lookout())

	linkedList.Reverse()

	fmt.Printf("%v\n", linkedList.Lookout())

	fmt.Println("-------------------------")


}

func TestSLinkedList() {
	linkedList := SLinkedList{}

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