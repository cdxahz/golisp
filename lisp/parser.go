package lisp

import (
	"fmt"
)

func Parse(tokens []Token) *Node {
	var current, root *Node
	count := 0
	current = new(Node)
	root = current
	for _, token := range tokens {

		if token.Type == OP || isAssign(token) {
			current.root = token
		} else if token.Type == NUMBER || token.Type == LEFT || token.Type == WORD {
			if count == 0 {
				continue
			}
			child := new(Node)
			if current.left == nil {
				current.left = child
			} else if current.right == nil {
				current.right = child
			} else {
				newRoot := new(Node)
				newRoot.left = current
				if current == root {
					root = newRoot
				}
				if current.parent != nil {
					if current == current.parent.left {
						current.parent.left = newRoot
					} else if current == current.parent.right {
						current.parent.right = newRoot
					}
					newRoot.parent = current.parent
					current.parent = newRoot
				}
				//copy the op
				newRoot.root = current.root

				current = newRoot
				current.right = child
			}
			if token.Type == NUMBER {
				child.root = token
			}
			child.parent = current
			if token.Type == LEFT {
				current = child
			}
			if token.Type == WORD {
				child.root = token
			}
		} else if token.Type == RIGHT {
			current = current.parent
		} else {
			DEBUG("current token : ", token)
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
	} else if isAssign(ast.root) {
		fmt.Println("=", string(ast.left.root.Value), string(ast.right.root.Value), result)
	}
	return string(ast.root.Value)
}

func isAssign(token Token) bool {
	return token.Type == WORD && string(token.Value) == "setf"
}
