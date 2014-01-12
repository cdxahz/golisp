package lisp

import(
    "testing"
    "fmt"
)

func TestScan(t *testing.T){

    source := "(* 10 (+ 1 2.3))"
    scanner := NewScanner(source)
	for _, token := range scanner.Tokens(){
		fmt.Println(token.ToString())
	}
}
