package main

import (
	"./lisp"
	"fmt"
    "os"
	"flag"
	"bufio"
    "strings"
)

var file = flag.String("file", "sample.lisp", "help message for file")

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
    fmt.Println(lisp.Eval(exp))

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
        parser = lisp.NewParser(scanner.Tokens())
        if strings.Contains(string(line), "defun"){
            ast := parser.MatchFunction()
            fmt.Println(string(line), " => \n function :", ast.Name, ", args:", ast.Args, ", tree:")
            lisp.Visit(ast.Expression, 0)
        }else{
		    ast,_ := parser.MatchExpression()
		    fmt.Println(string(line), " = ", lisp.Eval(ast))
        }
	}

	lisp.PrintTable()
    
}
