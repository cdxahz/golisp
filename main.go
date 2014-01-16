package main

import(
	"fmt"
	"os"
	"flag"
	"bufio"

	"github.com/cdxahz/golisp/lisp"
)

var file = flag.String("file", "sample.lisp", "help message for file")

func main(){

	flag.Parse()
	f, err := os.Open(*file)
	defer f.Close()

	if err != nil{
		panic(err)
	}
	
	reader := bufio.NewReader(f)
	for {
		line, _, _ := reader.ReadLine()
		if len(line) <= 0{
			break
		}
		scanner := lisp.NewScanner(string(line))
		ast := lisp.Parse(scanner.Tokens())
		fmt.Println(string(line), " = ", lisp.Eval(ast))

	}

	lisp.PrintTable()

}
