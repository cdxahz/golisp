package lisp

import(
	"fmt"
)

var table map[string] interface{}

func Eval(ast *Node) int {
	var op []byte 
	var left, right, result int
	if ast == nil {
		panic("invalid ast")
	}
	op = ast.root.Value
	if ast.root.Type == OP {
		if ast.left != nil {
			if ast.left.root.Type == NUMBER {
				left = toInt(ast.left.root.Value)
			} else if ast.left.root.Type == OP {
				left = Eval(ast.left)
			} else {
				panic("not supported token type")
			}
		}
		if ast.right != nil {
			if ast.right.root.Type == NUMBER {
				right = toInt(ast.right.root.Value)
			} else if ast.right.root.Type == OP {
				right = Eval(ast.right)
			} else {
				panic("not supported token type")
			}
		}

		if var_i == 0 {
			var_i = '0'
		}
		result = eval(string(op), left, right)
		return result
	}else if isAssign(ast.root){
		AddToTable(string(ast.left.root.Value), toInt(ast.right.root.Value))
		return toInt(ast.right.root.Value)
	}
	return toInt(ast.root.Value)
}

func eval(op string, left, right int) int{
	switch{
	case op == "+":
		return left + right
	case op == "-":
		return left - right
	case op == "*":
		return left * right
	case op == "/":
		return left / right
	}

	panic("not supported op " + string(op))
}

func toInt(val []byte) int{
	result := 0
	for _, v := range val{
		factor := 10
		result = result * factor + int(v - '0')
	}
	return int(result)
}

func AddToTable(key string, v interface{}) {
	if table == nil{
		table = make(map[string] interface{})	
	}

	table[key] = v
}

func PrintTable() {
	fmt.Println("[print the table info]:")
	if table != nil{
		for key, value := range table{
			fmt.Println(key, " = ", value)
		}
	}
}
