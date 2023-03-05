// Implementation of Edsger Dijkstra's the "shunting yard" algorithm
// https://en.wikipedia.org/wiki/Shunting_yard_algorithm

package parenthesis

import (
	"github.com/Goathy/stack"
)

const (
	opPow   string = "^"
	opMulti string = "*"
	opDiv   string = "/"
	opAdd   string = "+"
	opSub   string = "-"

	opLeftPar  string = "("
	opRightPar string = ")"
)

type associativity int

const (
	assocLeft  associativity = -1
	assocRight associativity = 1
)

func assoc(o string) associativity {
	switch o {
	case "^":
		return assocRight
	default:
		return assocLeft
	}
}

func precedence(op string) int {
	switch op {
	case "^":
		return 4
	case "*", "/":
		return 3
	case "+", "-":
		return 2
	default:
		return 0
	}
}

func Postfix(infix []string) []string {
	var (
		size    = int64(len(infix))
		ops, _  = stack.New[string](size)
		postfix = make([]string, 0)
	)

	for _, token := range infix {
		switch token {
		case opLeftPar:
			ops.Push(token)
		case opRightPar:
			for {
				operator, _ := ops.Pop()

				if operator == opLeftPar {
					break
				}

				postfix = append(postfix, operator)
			}
		case opAdd, opSub, opMulti, opDiv, opPow:
			for operator, _ := ops.Peek(); !ops.IsEmpty() && precedence(operator) > precedence(token) || precedence(operator) == precedence(token) && assoc(token) == assocLeft; operator, _ = ops.Peek() {
				operator, _ = ops.Pop()
				postfix = append(postfix, operator)
			}
			ops.Push(token)
		default:
			postfix = append(postfix, token)
		}
	}

	for !ops.IsEmpty() {
		operator, _ := ops.Pop()
		postfix = append(postfix, operator)
	}

	return postfix
}
