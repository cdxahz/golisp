package lisp

import (
	"fmt"
)

type Number interface{}

func isAllInt(a, b Number) bool {
	_, ok_a := a.(int)
	_, ok_b := b.(int)
	return ok_a && ok_b
}

func intValue(a, b Number) (int, int) {
	val_a, _ := a.(int)
	val_b, _ := b.(int)
	return val_a, val_b
}

func float64Value(a, b Number) (float64, float64) {
	val_ai, ok_ai := a.(int)
	val_bi, ok_bi := b.(int)

	val_a, _ := a.(float64)
	val_b, _ := b.(float64)
	if ok_ai {
		if ok_bi {
			return float64(val_ai), float64(val_bi)
		} else {
			return float64(val_ai), val_b
		}
	} else {
		if ok_bi {
			return val_a, float64(val_bi)
		} else {
			return val_a, val_b
		}
	}
}

var table map[string]interface{}

func Eval(ast *Node) Number {
	var op []byte
	var result, result_tmp Number
	if ast == nil {
		return ""
	} else if len(ast.Targets) == 0 {
		return 0
	}

	if ast.Op.Type == OP {
		op = ast.Op.Value
		for i := 0; i < len(ast.Targets); i++ {
			child := ast.Targets[i]
			if i == 0 {
				if child.Op.Type == NUMBER {
					result = toNumber(child.Op.Value)
				} else {
					result = Eval(&child)
				}
				continue
			}

			if child.Op.Value != nil && child.Op.Type == NUMBER {
				result_tmp = toNumber(child.Op.Value)

			} else if child.Op.Type == OP {
				result_tmp = Eval(&child)
			} else {
				panic("not supported token type")
			}
			result = eval(string(op), result, result_tmp)
		}

		return result
	} else {
		//tmp process
		AddToTable(string(ast.Targets[0].Op.Value), toNumber(ast.Targets[1].Op.Value))
		return toNumber(ast.Targets[1].Op.Value)
	}
	return toNumber(ast.Op.Value)
}

func eval(op string, left, right Number) Number {

	if isAllInt(left, right) {
		v_left, v_right := intValue(left, right)
		switch {
		case op == "+":
			return v_left + v_right
		case op == "-":
			return v_left - v_right
		case op == "*":
			return v_left * v_right
		case op == "/":
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

func toFloat(val []byte) float64 {
	result := 0.0
	factor := 10.0
	for _, v := range val {
		if !isDot(v) {
			if factor == 10.0 {
				result = result*factor + float64(v-'0')
			} else {
				result = result + float64(v-'0')*factor
			}
		} else {
			factor = 0.1
		}

	}
	return float64(result)
}

func toNumber(val []byte) Number {
	if isInt(val) {
		return toInt(val)
	} else if isFloat(val) {
		return toFloat(val)
	} else {
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
