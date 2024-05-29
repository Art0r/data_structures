package linkedlist

import "fmt"

type DLinkedList struct {
	head  *Node
	tail  *Node
	count int32
}

func (list *DLinkedList) Append(value int) {
	newNode := &Node{data: value}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		list.count++
		return
	}

	newNode.previous = list.tail
	list.tail.next = newNode
	list.tail = newNode
	list.count++
}

func (list *DLinkedList) Prepend(value int) {
	newNode := &Node{data: value}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
		list.count++
		return
	}

	newNode.next = list.head
	list.head.previous = newNode
	list.head = newNode
	list.count++
}

func (list *DLinkedList) Insert(index int, value int) {

	if index >= int(list.count) || index < 0 {
		return
	}

	newNode := &Node{data: value}

	if index == 0 {

		if list.head == nil {
			list.head = newNode
			list.tail = newNode
			list.count++
			return
		}
		
		newNode.next = list.head
		list.head.previous = newNode
		list.head = newNode
		list.count++
		return
	}

	if index == int(list.count) {
		list.tail.next = newNode
		newNode.previous = list.tail
		list.tail = newNode
		list.count++
		return
	}

	leader := list.traverseToIndex(index)
	holder := leader.next

	leader.next = newNode
	newNode.previous = leader
	newNode.next = holder
	if holder != nil {
		holder.previous = newNode
	}

	list.count++
}

func (list *DLinkedList) Remove(index int) {
	if index >= int(list.count) || index < 0 {
		return
	}

	if index == 0 {
		
		if list.head == nil {
			return
		}

		list.head = list.head.next

		if list.head != nil {
			list.head.previous = nil
			list.count--
			return
		}

		list.tail = nil
		list.count--
		return
	}

	leader := list.traverseToIndex(index)
	toRemove := leader.next

	if toRemove == list.tail {
		list.tail = leader
		leader.next = nil
		list.count--
		return
	}

	leader.next = toRemove.next
	if toRemove.next != nil {
		toRemove.next.previous = leader
	}

	list.count--
}

func (list *DLinkedList) FrontLookout() []int {
	currentNode := list.head

	finalList := []int{}
	for currentNode != nil {
		finalList = append(finalList, currentNode.data)
		currentNode = currentNode.next
	}

	return finalList
}

func (list *DLinkedList) BackLookout() []int {
	currentNode := list.tail

	finalList := []int{}
	for currentNode != nil {
		finalList = append(finalList, currentNode.data)
		currentNode = currentNode.previous
	}

	return finalList
}

func (list *DLinkedList) traverseToIndex(index int) *Node {
	currentNode := list.head
	for i := 0; i < index-1; i++ {
		if currentNode == nil {
			break
		}
		currentNode = currentNode.next
	}
	return currentNode
}

func TestDLinkedList() {
	linkedList := DLinkedList{}

	linkedList.Append(1)
	linkedList.Append(2)
	linkedList.Append(3)
	linkedList.Append(4)

	fmt.Println("-------------------------")

	fmt.Printf("%v\n", linkedList.FrontLookout())
	fmt.Printf("%v\n", linkedList.BackLookout())
	fmt.Println("HEAD: ", linkedList.head.data)
	fmt.Println("TAIL: ", linkedList.tail.data)
	fmt.Println("TOTAL: ", linkedList.count)
	fmt.Println("-------------------------")

	linkedList.Insert(2, 5)

	fmt.Printf("%v\n", linkedList.FrontLookout())
	fmt.Printf("%v\n", linkedList.BackLookout())
	fmt.Println("HEAD: ", linkedList.head.data)
	fmt.Println("TAIL: ", linkedList.tail.data)
	fmt.Println("TOTAL: ", linkedList.count)
	fmt.Println("-------------------------")

	linkedList.Remove(3)

	fmt.Printf("%v\n", linkedList.FrontLookout())
	fmt.Printf("%v\n", linkedList.BackLookout())
	fmt.Println("HEAD: ", linkedList.head.data)
	fmt.Println("TAIL: ", linkedList.tail.data)
	fmt.Println("TOTAL: ", linkedList.count)
	fmt.Println("-------------------------")

	linkedList.Remove(3)

	fmt.Printf("%v\n", linkedList.FrontLookout())
	fmt.Printf("%v\n", linkedList.BackLookout())
	fmt.Println("HEAD: ", linkedList.head.data)
	fmt.Println("TAIL: ", linkedList.tail.data)
	fmt.Println("TOTAL: ", linkedList.count)
	fmt.Println("-------------------------")

	linkedList.Append(2)
	linkedList.Append(50)

	fmt.Printf("%v\n", linkedList.FrontLookout())
	fmt.Printf("%v\n", linkedList.BackLookout())
	fmt.Println("HEAD: ", linkedList.head.data)
	fmt.Println("TAIL: ", linkedList.tail.data)
	fmt.Println("TOTAL: ", linkedList.count)
	fmt.Println("-------------------------")

	linkedList.Prepend(7)
	linkedList.Prepend(12)

	fmt.Printf("%v\n", linkedList.FrontLookout())
	fmt.Printf("%v\n", linkedList.BackLookout())
	fmt.Println("HEAD: ", linkedList.head.data)
	fmt.Println("TAIL: ", linkedList.tail.data)
	fmt.Println("TOTAL: ", linkedList.count)
	fmt.Println("-------------------------")

}
