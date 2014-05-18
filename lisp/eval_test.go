package lisp

import (
	"fmt"
	"testing"
)

func TestEval(t *testing.T) {
	source := "(* 5 (- 10 (+ 1 2)))"

	scanner := NewScanner(source)
	parser := NewParser(scanner.Tokens())
	ast := parser.Parse()

	fmt.Println("test eval :", Eval(ast))

	if Eval(ast) != 35 {
		t.Fail()
	}
}
