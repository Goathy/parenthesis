package parenthesis

import (
	"strings"

	"github.com/Goathy/stack"
)

func precedence(op rune) int {
	switch op {
	case '^':
		return 3
	case '*', '/':
		return 2
	case '+', '-':
		return 1
	default:
		return 0
	}
}

func Postfix(infix string) string {
	op, _ := stack.New[rune](-1)
	postfix := strings.Builder{}

	for _, in := range infix {
		switch in {
		case '(':
			op.Push(in)
		case ')':
			for par, _ := op.Pop(); par != '('; par, _ = op.Pop() {
				postfix.WriteRune(par)
			}
		case '+', '-', '*', '/', '^':
			for {
				x, _ := op.Peek()
				if !op.IsEmpty() && precedence(x) >= precedence(in) {
					x, _ = op.Pop()
					postfix.WriteRune(x)
				} else {
					break
				}
			}
			op.Push(in)
		default:
			postfix.WriteRune(in)
		}
	}

	for !op.IsEmpty() {
		op, _ := op.Pop()
		postfix.WriteRune(op)
	}

	return postfix.String()
}
