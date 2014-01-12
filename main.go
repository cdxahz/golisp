package main

import(
	"fmt"
	"github.com/cdxahz/golisp/lisp"
)

func main(){
	var source string
	fmt.Println("Please input the expression:")
	fmt.Scan(source)
	scanner := lisp.NewScanner(source)
	ast := lisp.Parse(scanner.Tokens())
	fmt.Println(lisp.Eval(ast))

}
