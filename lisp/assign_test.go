package lisp

import (
	"fmt"
	"testing"
)

func TestAssign(t *testing.T) {
	source := "(setf abc 123)"

	scanner := NewScanner(source)
	ast := Parse(scanner.Tokens())

	fmt.Println("test setf:", ast.root.ToString(), ast.left.root.ToString())
	fmt.Println("test setf:", Gen(ast))
	fmt.Println("test eval setf:", Eval(ast))

	if string(ast.root.Value) != "setf" {
		t.Fail()
	}

	if string(ast.left.root.Value) != "abc" {
		t.Fail()
	}
}
