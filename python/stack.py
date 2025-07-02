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


class Stack:
    _top: Optional[_Node]
    _bottom: Optional[_Node]

    def __init__(self):
        self._top = None
        self._bottom = None

    @property
    def top(self):
        return self._top.value

    @property
    def bottom(self):
        return self._bottom.value

    def is_empty(self):
        return self._top is None

    def pop(self):

        if self.is_empty():
            return

        self._top = self._top.next
        self.print_stack_with_top_and_bottom()

    def push(self, value):
        temp = _Node(value)

        if self.is_empty():
            self._bottom = temp
            self._top = temp
            self.print_stack_with_top_and_bottom()
            return

        previous_top = self._top
        self._top = temp
        self._top.next = previous_top
        self.print_stack_with_top_and_bottom()

    def print_stack(self):
        current = self._top

        text = ""
        while current is not None:
            text += "{0}".format(current.value)

            if current.next is not None:
                text += "\n"

            current = current.next

        print(text)

    def print_stack_with_top_and_bottom(self):
        self.print_stack()
        print("TOP:", self.top)
        print("BOTTOM:", self.bottom)
        print("---------------------------")


stack = Stack()

stack.push(1)
stack.push(2)
stack.push(3)
stack.push(4)
stack.push(5)

stack.pop()
stack.pop()
