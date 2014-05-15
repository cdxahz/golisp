package lisp

import(
    "testing"
)

func TestParse(t *testing.T){
    scanner := NewScanner("file name")
    tokens := scanner.Scan()
    parser := NewParser(tokens)
    parser.parse()
    parser.evaluate()
}
