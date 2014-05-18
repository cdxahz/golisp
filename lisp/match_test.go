package lisp

import (
	"../lisp"
	"log"
	"testing"
)

func TestMatchExpression(t *testing.T) {
	scanner := lisp.NewScanner("(+ a b)")
	tokens := scanner.Tokens()
	for _, token := range tokens {
		log.Println(token.ToString())
	}

	parser := lisp.NewParser(tokens)

	node := parser.MatchExpression()

	log.Println(node)

	panic("show")
}
