package lisp

import (
	"log"
	"testing"
)

func TestParseFunction(t *testing.T) {
	scanner := NewScanner("(defun add (a b) (+ a b))")
	tokens := scanner.Tokens()
	log.Println(tokens)
	for _, token := range tokens {
		log.Println(token.ToString())
	}
	parser := NewParser(tokens)
	function := parser.MatchFunction()
	if function == nil {
		t.Fail()
	}
	log.Println(function.Name)
	for _, arg := range function.Args {
		log.Println(arg.ToString())
	}
	visit(function.Expression)

}

func visit(node *Node) {
	if node == nil {
		return
	}
	log.Println("node -> ", node.Op.ToString())
	for _, n := range node.Targets {
		visit(&n)
	}
}
