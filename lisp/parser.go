package lisp

import (
	"fmt"
)

type Parser struct {
	tokens  []Token
	ast     *Node
	current int
}

type Node struct {
	Op      Token
	Targets []Node
}

type Function struct {
	Name       string
	Args       []Token
	Expression *Node
}

func NewParser(tokens []Token) *Parser {
	return &Parser{
		tokens,
		nil,
		-1,
	}
}

func (parser *Parser) nextToken() (Token, bool) {
	parser.current = parser.current + 1
	if parser.current >= 0 && parser.current <= len(parser.tokens)-1 {
		token := parser.tokens[parser.current]
		return token, true
	} else {
		return Token{}, false
	}
}

func (parser *Parser) MatchFunction() *Function {
	if _, ok := parser.match("("); ok {
		if _, ok_fun := parser.match("defun"); ok_fun {
			functionName := string(parser.matchValue())
			args := parser.matchArgs()
			expression, _ := parser.MatchExpression()
			parser.match(")")
			function := &Function{
				functionName,
				args,
				expression,
			}
			AddToTable(functionName, function)
			return function
		}
	} else {
		panic("invalid function")
	}

	panic("parse function failed")
}

func (parser *Parser) match(toMatch string) (Token, bool) {
	token, _ := parser.nextToken()
	return token, string(token.Value) == toMatch
}

func (parser *Parser) matchValue() []byte {
	token, _ := parser.nextToken()
	return token.Value
}

func (parser *Parser) currentToken() Token {
	return parser.tokens[parser.current]
}

func (parser *Parser) matchArgs() []Token {
	parser.match("(")
	nodes := make([]Token, 0)
	for {
		token, _ := parser.nextToken()
		if string(token.Value) != ")" {
			nodes = append(nodes, token)
		} else {
			break
		}
	}
	return nodes
}

func (parser *Parser) Parse() *Node {
	if node, ok := parser.MatchExpression(); ok {
		return node
	}
	return nil
}

func (parser *Parser) MatchExpression() (*Node, bool) {
	stepOne := parser.lookForward()
	if string(stepOne.Value) == "(" {
		parser.match("(")
		var op Token
		if parser.lookForward().Type == OP {
			op = parser.matchOp()
		} else {
			op, _ = parser.nextToken()
		}
		return &Node{
			op,
			parser.MatchExpressions(),
		}, true
	} else if stepOne.Type == NUMBER || stepOne.Type == WORD {
		token, _ := parser.nextToken()
		return &Node{
			token,
			[]Node{},
		}, true
	} else if stepOne.Type == RIGHT {
		parser.nextToken()
		return &Node{}, false
	} else {
		panic("should not touch here")
	}
	panic("invalid expression match")
}

func (parser *Parser) MatchExpressions() []Node {
	expressions := make([]Node, 0)
	for {
		if parser.lookForward().Value == nil {
			return expressions
		}
		expression, ok := parser.MatchExpression()
		if ok {
			expressions = append(expressions, *expression)
		} else {
			break
		}
	}
	return expressions
}

func (parser *Parser) lookForward() Token {
	if parser.current >= len(parser.tokens)-1 {
		return Token{}
	}
	var token Token
	token = parser.tokens[parser.current+1]
	return token
}

func (parser *Parser) matchType(toMatch int) (Token, bool) {
	token, _ := parser.nextToken()
	if token.Type == toMatch {
		return token, true
	} else {
		panic("invalid type match")
	}
}

func (parser *Parser) matchOp() Token {
	if token, ok := parser.matchType(OP); ok {
		return token
	}
	panic("invalid op match")
}

func Visit(node *Node, depth int) {
	if node != nil && node.Op.Value != nil {
		for i := 0; i < depth; i++ {
			fmt.Print("\t")
		}
		fmt.Println("=>", node.Op.ToString())
	}

	for _, item := range node.Targets {
		Visit(&item, depth+1)
	}
}

func DEBUG(msg string, v ...interface{}) {
	fmt.Println(msg, v)
}

var var_i byte

func Gen(ast *Node) string {
	var op, result string
	if ast == nil {
		return ""
	} else if len(ast.Targets) == 0 {
		return string(ast.Op.Value)
	}

	first := ast.Targets[0]

	if ast.Op.Type == OP {
		op = string(ast.Op.Value)
		for i := 0; i < len(ast.Targets); i++ {
			if i == 0 {
				result = string(first.Op.Value)
				continue
			}
			child := ast.Targets[i]
			if var_i == 0 {
				var_i = '0'
			}
			var_i = var_i + 1
			result_tmp := "_var" + string(var_i)

			if child.Op.Value != nil {
				fmt.Println(op, result, Gen(&child), result_tmp)
				result = result_tmp
			}
		}
		return result
	} else {
		var_i = var_i + 1
	}
	return string(ast.Op.Value)
}

func isAssign(token Token) bool {
	return token.Type == WORD && string(token.Value) == "setf"
}
