package lisp

import (
	"fmt"
)

type Node struct {
	root   Token
	left   *Node
	right  *Node
	name   string
	parent *Node
}

func Parse(tokens []Token) *Node {
	var current, root *Node
	count := 0
	current = new(Node)
	root = current
	for _, token := range tokens {

		if token.Type == OP {
			current.root = token
		} else if token.Type == NUMBER || token.Type == LEFT {
			if count == 0 {
				continue
			}
			child := new(Node)
			if current.left == nil {
				current.left = child
			} else {
				current.right = child
			}
			if token.Type == NUMBER {
				child.root = token
			}
			child.parent = current
			if token.Type == LEFT {
				current = child
			}
		} else if token.Type == RIGHT {
			current = current.parent
		} else {
			panic("not support")
		}
		count = count + 1
		//DEBUG("parsing ", token, current, root)
	}
	return root
}

func PrintAst(root *Node) {
	if root == nil {
		return
	}

	fmt.Println(root.root.ToString())
	PrintAst(root.left)
	PrintAst(root.right)
}

func DEBUG(msg string, v ...interface{}) {
	fmt.Println(msg, v)
}

var var_i byte

func Gen(ast *Node) string {
	var op, left, right, result string
	if ast == nil {
		return ""
	}
	if ast.root.Type == OP {
		op = string(ast.root.Value)
		if ast.left != nil {
			if ast.left.root.Type == NUMBER {
				left = string(ast.left.root.Value)
			} else if ast.left.root.Type == OP {
				left = Gen(ast.left)
			} else {
				panic("not supported token type")
			}
		}
		if ast.right != nil {
			if ast.right.root.Type == NUMBER {
				right = string(ast.right.root.Value)
			} else if ast.right.root.Type == OP {
				right = Gen(ast.right)
			} else {
				panic("not supported token type")
			}
		}

		if var_i == 0 {
			var_i = '0'
		}
		var_i = var_i + 1
		result = "_var" + string(var_i)
		fmt.Println(op, left, right, result)
		return result
	}
	return string(ast.root.Value)
}
