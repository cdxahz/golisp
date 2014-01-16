package lisp

import(
	"testing"
	"fmt"
)

func TestEval(t *testing.T){
	source := "(* 5 (- 10 (+ 1 2)))"

	scanner := NewScanner(source)
	ast := Parse(scanner.Tokens())

	fmt.Println("test eval :", Eval(ast))

	if Eval(ast) != 35{
		t.Fail()
	}
}
