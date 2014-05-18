package lisp

import (
	"fmt"
	"testing"
)

func TestAssign(t *testing.T) {
	source := "(setf abc 123)"

	scanner := NewScanner(source)
    parser := NewParser(scanner.Tokens())
	ast := parser.Parse()

	fmt.Println("test setf:", Gen(ast))
	fmt.Println("test eval setf:", Eval(ast))

	if string(ast.Op.Value) != "setf" {
		t.Fail()
	}

}
