package stacks

// Stack
type stack []int

// Empty returns true when stack does not have elements
func (s stack) Empty() bool {
	return len(s) == 0
}

func (s stack) Clear() {
	for !s.Empty() {
		s.Pop()
	}
}

// Push adds an element into the stack
func (s stack) Push(v int) stack {
	return append(s,v)
}

// Pop removes the last element pushed from the stack
// Returns false when stack is empty
func (s stack) Pop() (stack, bool) {
	if s.Empty() {
		return s, false
	}

	length := len(s)
	return s[:length-1], true
}

// Top returns the last element pushed into the stack
// Returns false when the stack is empty
func (s stack) Top() (int, bool) {
	if s.Empty() {
		return -1, false
	}

	length := len(s)
	return s[length-1], true
}

func (s stack) Size() (int) {
	return len(s);
}