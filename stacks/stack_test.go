package stacks

import "testing"

func TestStack(t *testing.T) {
	s := make(stack,0)

	s, ok := s.Pop()
	if ok != false {
		t.Fatalf("Pop in empty container should return false")
	}

	// Push a 1
	s = s.Push(1)
	res, ok := s.Top()
	if (!ok || res != 1 || s.Size() != 1) {
		t.Log("Res:",res, "Ok:", ok)
		t.Fatalf("Pushing 1 into stack is not working")
	}

	// Push a 2
	s = s.Push(2)
	res, ok = s.Top()
	if (!ok || res != 2 || s.Size() != 2) {
		t.Log("Res:",res, "Ok:", ok)
		t.Fatalf("Pushing 2 into stack is not working")
	}

	// Pop element from container
	s, ok = s.Pop()
	if (ok == false && s.Size()!=1) {
		t.Log("Res:",res, "Ok:", ok)
		t.Fatalf("Pop should pass in non empty container")
	}
	res, ok = s.Top()
	if (!ok || res != 1 || s.Size() != 1) {
		t.Log("Res:",res, "Ok:", ok)
		t.Fatalf("Top element should be 1")
	}

	// Push element and print top
	s = s.Push(3)
	res, ok = s.Top()
	if (!ok || res != 3 || s.Size() != 2) {
		t.Log("Res:",res, "Ok:", ok)
		t.Fatalf("Pushing 3 into stack is not working")
	}
}