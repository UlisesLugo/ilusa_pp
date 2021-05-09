package stacks

// Stack
type Stack []int

// Empty returns true when Stack does not have elements
func (s Stack) Empty() bool {
	return len(s) == 0
}

func (s Stack) Clear() {
	for !s.Empty() {
		s.Pop()
	}
}

// Push adds an element into the Stack
func (s Stack) Push(v int) Stack {
	return append(s,v)
}

// Pop removes the last element pushed from the Stack
// Returns false when Stack is empty
func (s Stack) Pop() (Stack, bool) {
	if s.Empty() {
		return s, false
	}

	length := len(s)
	return s[:length-1], true
}

// Top returns the last element pushed into the Stack
// Returns false when the Stack is empty
func (s Stack) Top() (int, bool) {
	if s.Empty() {
		return -1, false
	}

	length := len(s)
	return s[length-1], true
}

func (s Stack) Size() (int) {
	return len(s);
}