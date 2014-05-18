package main

import (
	"./lisp"
	"fmt"
)

func main() {
	scanner := lisp.NewScanner("(defun add (a b c) (+ a b c (- a 1) d))")
	tokens := scanner.Tokens()
	parser := lisp.NewParser(tokens)
	node := parser.MatchFunction()
	fmt.Print("name : ", node.Name, "-> args :")
	for _, arg := range node.Args {
		fmt.Print(arg.ToString(), " ")
	}
	fmt.Println()
	lisp.Visit(node.Expression, 0)
	lisp.Gen(node.Expression)

	expression := lisp.NewScanner("(- 2 1.0 (+ 3 4) 5)").Tokens()
	parser_exp := lisp.NewParser(expression)
	exp, _ := parser_exp.MatchExpression()
	lisp.Gen(exp)
}
