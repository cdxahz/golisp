package lisp

type Parser struct{
    tokens []Token
    ast *Node
    current int
}

const(
    MAX_ARGS = 10
)

type Node struct{
    Op Token
    Targets []Node
}

type Function struct{
    name string
    args []Token
    expression *Node
}

func NewParser(tokens []Token) *Parser{
    return &Parser{
        tokens,
        nil,
        0,
    }
}

func (parser *Parser) nextToken() Token{
   parser.current = parser.current + 1
   return parser.tokens[parser.current]
}

func (parser *Parser) matchFunction() *Function{
    if parser.match("(") && parser.match("defun"){
        functionName := string(parser.matchValue())
        args := parser.matchArgs()
        expression := parser.matchExpression()
        parser.match(")")

        return &Function{
            functionName,
            args,
            expression,
        }
    }else{
        panic("invalid function")
    }
}

func (parser *Parser) match(toMatch string) bool{
    token := parser.nextToken()
    return string(token.Value) == toMatch
}

func (parser *Parser) matchValue() []byte{
    return parser.nextToken().Value
}

func (parser *Parser) currentToken() Token{
    return parser.tokens[parser.current]
}

func (parser *Parser) matchArgs() []Token{
    parser.match("(")
    nodes := make([]Token, MAX_ARGS)
    for{
        token := parser.nextToken()
        if string(token.Value) != ")"{
            nodes = append(nodes, token)
        }else{
            break
        }
    }
    return nodes
}

func (parser *Parser) matchExpression() *Node{
    tokens := make([]Node, MAX_ARGS)
    stepOne := parser.lookForward()
    if string(stepOne.Value) == "("{
        parser.match("(")
        if parser.lookForward().Type == OP{
            return &Node{
                parser.matchOp(),
                parser.matchExpressions(),
            }
        }else if parser.lookForward().Type == WORD{
            return  &Node{
                Token{},
                append(tokens, Node{parser.nextToken(), []Node{},}),
            }
        }else{
            parser.match(")")
        }
    }else{
        return &Node{
            Token{[]byte("="), OP,},
            append(tokens, Node{parser.nextToken(), []Node{}}),
        }
    }
    panic("invalid expression match")
}

func (parser *Parser) matchExpressions() []Node{
    expressions := make([]Node, 0)
    for {
        if parser.lookForward().Value == nil{
            return expressions
        }
        expressions = append(expressions, *parser.matchExpression())
    }
    return expressions
}

func (parser *Parser) lookForward() Token{
    return parser.tokens[parser.current + 1] 
}

func (parser *Parser) matchOp() Token{
    token := parser.nextToken()
    if token.Type == OP {
        return token
    }else{
        panic("invalid op")
    }
}
