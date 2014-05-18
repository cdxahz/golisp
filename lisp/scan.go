package lisp

import (
	"fmt"
	"strings"
)

const (
	NUMBER = iota
	WORD
	OP
	LEFT
	RIGHT
	INVALID
)

const (
	END = iota
)

type Token struct {
	Value []byte
	Type  int
}

type Scanner struct {
	src     []byte
	pos     int
	lastPos int
}

func NewScanner(source string) *Scanner {
	buffer := strings.NewReader(source)
	var src []byte = make([]byte, buffer.Len())
	n, err := buffer.Read(src)
	if err != nil {
		fmt.Println("size :", n)
		panic(err)
	}
	return &Scanner{src, 0, 0}
}

func (scanner *Scanner) Current() byte {
	//fmt.Printf("last pos: %v  current pos: %v buffer size : %v", scanner.lastPos, scanner.pos, len(scanner.src))
	if len(scanner.src) <= 0 || scanner.pos >= len(scanner.src) {
		//	fmt.Println("current: END")
		return 0
	}
	//fmt.Printf(" current: %v\n", string(scanner.src[scanner.pos]))
	return scanner.src[scanner.pos]
}

func (scanner *Scanner) LookAhead() byte {
	if scanner.pos >= len(scanner.src)-1 {
		return END
	}
	return scanner.src[scanner.pos+1]
}

func (scanner *Scanner) Tokens() []Token {
	tokens := []Token{}
	for {
		tok := scanner.NextToken()
		if tok.Type == INVALID {
			break
		}
		tokens = append(tokens, *tok)
	}
	return tokens
}

func (scanner *Scanner) NextToken() *Token {

	for {
		current := scanner.Current()
		if isDigital(current) {
			if !isDigital(scanner.LookAhead()) && !isDot(scanner.LookAhead()) {
				return scanner.next(NUMBER)
			}
		} else if isOp(current) {
			if !isOp(scanner.LookAhead()) {
				return scanner.next(OP)
			}
		} else if isAlpha(current) {
			if !isAlpha(scanner.LookAhead()) {
				return scanner.next(WORD)
			}
		} else if isDot(current) {
			scanner.GoAhead()
			continue
		} else if isBlank(current) {
			scanner.skip()
			continue
		} else if isLeftMark(current) {
			return scanner.next(LEFT)
		} else if isRightMark(current) {
			return scanner.next(RIGHT)
		} else if isEnd(current) {
			break
		} else {
			panic("not support")
		}

		scanner.GoAhead()

	}
	return &Token{nil, INVALID}
}

func (scanner *Scanner) GoAhead() byte {
	scanner.pos = scanner.pos + 1
	if scanner.pos >= len(scanner.src) {
		return 0
	}
	return scanner.src[scanner.pos]
}

func (scanner *Scanner) next(t int) *Token {

	token := &Token{scanner.src[scanner.lastPos : scanner.pos+1], t}
	scanner.GoAhead()
	scanner.lastPos = scanner.pos
	return token
}

func (scanner *Scanner) skip() {
	scanner.GoAhead()
	scanner.lastPos = scanner.lastPos + 1
}

func (token *Token) ToString() string {
	return fmt.Sprintf("token : %v -> type : %v", string(token.Value), token.Type)
}

func isDigital(ch byte) bool {
	return ch >= '0' && ch <= '9'
}

func isNumber(value []byte) bool {
	for _, ch := range value {
		if !isDigital(ch) && !isDot(ch) {
			return false
		}
	}
	return true
}

func isInt(value []byte) bool {
	for _, ch := range value {
		if !isDigital(ch) {
			return false
		}
	}
	return true
}

func isFloat(value []byte) bool {
	return isNumber(value) && !isInt(value)
}

func isOp(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func isAlpha(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z'
}

func isDot(ch byte) bool {
	return ch == '.'
}

func isBlank(ch byte) bool {
	return ch == ' '
}

func isLeftMark(ch byte) bool {
	return ch == '('
}

func isRightMark(ch byte) bool {
	return ch == ')'
}

func isEnd(ch byte) bool {
	return ch == END
}
