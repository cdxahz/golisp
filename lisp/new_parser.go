package lisp

type Parser struct{
    tokens []Token
    ast *Node
    current int
}

const(
    MAX_ARGS = 10
)

type Node interface{
    evaluate() interface{}
    toString() string
}

type Expression struct{
    op Token
    targets []Token
}

type Function struct{
    name string
    args []Token
    expression *Expression
}

func NewParser(tokens []Token) *Parser{
    return &Parser{
        tokens,
        nil,
        0,
    }
}

func (parser *Parser) parse(){

}

func (parser *Parser) evaluate(){

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

func (parser *Parser) matchExpression() *Expression{
    parser.match("(")
    op := parser.matchOp()
    tokens := make([]Token, MAX_ARGS)
    for{
        token := parser.nextToken()
        if string(token.Value) != ")"{
            tokens = append(tokens, token)
        }else{
            break
        }
    }
    return &Expression{
        op,
        tokens,
    }
}

func (parser *Parser) matchOp() Token{
    token := parser.nextToken()
    if token.Type == OP {
        return token
    }else{
        panic("invalid op")
    }
}
