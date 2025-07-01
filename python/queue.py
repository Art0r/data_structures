from typing import Optional


class _Node:
    _value: Optional[int]
    _next: Optional['_Node']

    def __init__(self, value):
        self._value = value
        self._next = None

    @property
    def value(self):
        return self._value

    @value.setter
    def value(self, new_value):
        self._value = new_value

    @property
    def next(self):
        return self._next

    @next.setter
    def next(self, new_next):
        self._next = new_next


class Queue:
    _front: Optional[_Node]
    _rear: Optional[_Node]

    def __init__(self):
        self._front = None
        self._rear = None

    def get_front(self):
        return self._front.value

    def get_rear(self):
        return self._rear.value

    def is_empty(self):
        return self._front is None

    def dequeue(self):

        if self.is_empty():
            return

        self._front = self._front.next
        self.print_queue_with_front_and_rear()

    def enqueue(self, value):
        temp = _Node(value)

        if self.is_empty():
            self._rear = temp
            self._front = temp
            self.print_queue()
            self.print_queue_with_front_and_rear()
            return

        self._rear.next = temp
        self._rear = temp
        self.print_queue_with_front_and_rear()

    def print_queue(self):
        current = self._front

        while current is not None:
            print(current.value)
            current = current.next

    def print_queue_with_front_and_rear(self):
        self.print_queue()
        print("FRONT:", self.get_front())
        print("REAR:", self.get_rear())
        print("---------------------------")


queue = Queue()

queue.enqueue(1)
queue.enqueue(2)
queue.enqueue(3)
queue.enqueue(4)
queue.enqueue(5)

queue.dequeue()
queue.dequeue()
