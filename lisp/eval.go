package lisp

import (
	"fmt"
)

type Number interface{}

func isAllInt(a, b Number) bool{
    _, ok_a := a.(int)
    _, ok_b := b.(int)
    return ok_a && ok_b
}

func intValue(a, b Number) (int, int){
    val_a, _ := a.(int)
    val_b, _ := b.(int)
    return val_a, val_b
}

func float64Value(a, b Number) (float64, float64){
    val_ai, ok_ai := a.(int)
    val_bi, ok_bi := b.(int)

    val_a, _ := a.(float64)
    val_b, _ := b.(float64)
    if ok_ai{
        if ok_bi{
            return float64(val_ai), float64(val_bi)
        }else{
            return float64(val_ai), val_b
        }
    }else{
        if ok_bi{
            return val_a, float64(val_bi)
        } else{
            return val_a, val_b
        }
    }
}

var table map[string]interface{}

func Eval(ast *Node) Number {
	var op []byte
	var left, right, result Number
	if ast == nil {
		panic("invalid ast")
	}
	op = ast.root.Value
	if ast.root.Type == OP {
		if ast.left != nil {
			if ast.left.root.Type == NUMBER {
				left = toNumber(ast.left.root.Value)
			} else if ast.left.root.Type == OP {
				left = Eval(ast.left)
			} else {
				panic("not supported token type")
			}
		}
		if ast.right != nil {
			if ast.right.root.Type == NUMBER {
				right = toNumber(ast.right.root.Value)
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
	} else if isAssign(ast.root) {
		AddToTable(string(ast.left.root.Value), toNumber(ast.right.root.Value))
		return toNumber(ast.right.root.Value)
	}
	return toNumber(ast.root.Value)
}

func eval(op string, left, right Number) Number {

    if isAllInt(left, right){
        v_left, v_right := intValue(left, right)
        switch{
        case op == "+":
            return v_left + v_right
        case op == "-":
            return v_left - v_right
        case op == "*":
            return v_left * v_right
        case op =="/":
            return v_left / v_right
        }
        panic("not supported")
    }
    v_leftf, v_rightf := float64Value(left, right)
	switch {
	case op == "+":
		return v_leftf + v_rightf
	case op == "-":
		return v_leftf - v_rightf
	case op == "*":
		return v_leftf * v_rightf
	case op == "/":
		return v_leftf / v_rightf
	}

	panic("not supported op " + string(op))
}

func toInt(val []byte) int {
	result := 0
	for _, v := range val {
		factor := 10
		result = result*factor + int(v-'0')
	}
	return int(result)
}

func toFloat(val []byte) float64{
    result := 0.0
    factor := 10.0
    for _, v :=range val{
        if !isDot(v){
            if factor == 10.0{
                result = result*factor + float64(v-'0')
            }else {
                result = result + float64(v-'0') * factor
            }  
        }else{
            factor = 0.1
        }

    }
    return float64(result)
}

func toNumber(val []byte) Number{
    if isInt(val){
        return toInt(val)
    }else if isFloat(val){
        return toFloat(val)
    }else{
        panic(string(val) + "not supported")
    }
}

func AddToTable(key string, v interface{}) {
	if table == nil {
		table = make(map[string]interface{})
	}

	table[key] = v
}

func PrintTable() {
	fmt.Println("[print the table info]:")
	if table != nil {
		for key, value := range table {
			fmt.Println(key, " = ", value)
		}
	}
}
