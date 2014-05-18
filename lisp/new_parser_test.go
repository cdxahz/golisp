package lisp

import(
    "testing"
    "log"
    "../lisp"
)

func TestParse(t *testing.T){
    scanner := lisp.NewScanner("(defun add (a b) (+ a b))")
    tokens := scanner.Tokens()
    log.Println(tokens)
    for _, token := range tokens {
        log.Println(token.ToString())
    }
    parser := lisp.NewParser(tokens)
    function := parser.MatchFunction()
    if function == nil{
        t.Fail()
    }
    log.Println(function.Name)
    for _, arg := range function.Args{
        log.Println(arg.ToString())
    }
    visit(function.Expression)

    panic("show")
}

func visit(node *lisp.Node){
    if node == nil{
        return
    }
    log.Println("node -> ", node.Op.ToString())
    for _, n := range node.Targets{
        visit(&n)
    }
}
