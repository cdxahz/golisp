package lisp

import (
	"fmt"
	"testing"
)

func TestScan(t *testing.T) {

	source := "(* 10 (+ 1 2.3))"
	scanner := NewScanner(source)
	count := 0
	for _, token := range scanner.Tokens() {
		fmt.Println(token.ToString())
		count = count + 1
	}

	if count != 9 {
		t.Fail()
	}

}
