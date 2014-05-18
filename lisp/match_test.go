package lisp

import (
	"log"
	"testing"
)

func TestMatchExpression(t *testing.T) {
	scanner := NewScanner("(+ a b)")
	tokens := scanner.Tokens()
	for _, token := range tokens {
		log.Println(token.ToString())
	}

	parser := NewParser(tokens)

	node, _ := parser.MatchExpression()

	log.Println(node)

}
