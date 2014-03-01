package lisp

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	stack := New()
	stack.Push(1)
	stack.Push("a")
	fmt.Println(stack.Top())
	stack.Pop()
	stack.Pop()
	if !stack.Empty() {
		t.Fail()
	}
}
