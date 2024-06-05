package queue


type Queue struct {
	Head    *Node
	Tail *Node
	Count  int32
}

func (queue *Queue) Enqueue(value int) {
	new_node := &Node{data: value}

	if queue.Head == nil {
		queue.Head = new_node
		queue.Tail = new_node
		queue.Count++
		return
	}

	new_node.next = queue.Head
	queue.Head = new_node
	queue.Count++
}

func (queue *Queue) Dequeue() {

	if queue.Head == nil || queue.Tail == nil {
		return
	}

	queue.Head = queue.Head.next

	queue.Count--
}

func (queue *Queue) Peek(index int) int {
	if queue.Count == 0 {
		return -1
	}

	node := queue.transverseStack(index)
	return node.data
}

func (queue *Queue) Lookout() []int {
	currentNode := queue.Head

	finalList := []int{}
	for currentNode != nil {
		finalList = append(finalList, currentNode.data)
		currentNode = currentNode.next
	}

	return finalList
}

func (queue *Queue) transverseStack(index int) *Node {
	currentNode := queue.Head

	for i := 0; i < index-1; i++ {
		if currentNode == nil {
			break
		}
		currentNode = currentNode.next
	}

	return currentNode
}