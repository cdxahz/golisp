package mylang

import(
    "testing"
    "fmt"
)

//FIX ME: the last token can't be parse
func TestScan(t *testing.T){

    source := "(* 10 (+ 1 2.3))"
    scanner := NewScanner(source)
	for _, token := range scanner.Tokens(){
		fmt.Println(token.ToString())
	}
}
