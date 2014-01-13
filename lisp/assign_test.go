package lisp

import(
	"testing"
	"fmt"
)

func TestAssign(t *testing.T){
	source := "(setf abc 123)"

	scanner := NewScanner(source)
	ast := Parse(scanner.Tokens())

	fmt.Println("test setf:", ast.root.ToString(), ast.left.root.ToString())
	fmt.Println("test setf:", Gen(ast))
	fmt.Println("test eval setf:", Eval(ast))
}
