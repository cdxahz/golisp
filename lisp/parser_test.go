package lisp

import (
	"fmt"
	"testing"
)

func TestParse(t *testing.T) {
	source := "(* (+ (- 10 1) 2) 10.1)"

	scanner := NewScanner(source)

	tokens := scanner.Tokens()
	for _, token := range tokens {
		fmt.Println(token.ToString())
	}

	ast := Parse(tokens)
	PrintAst(ast)
	fmt.Println()

	fmt.Println("gen the instructions")
	Gen(ast)

}
