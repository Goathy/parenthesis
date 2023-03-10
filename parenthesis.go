// Implementation of Edsger Dijkstra's the "shunting yard" algorithm
// https://en.wikipedia.org/wiki/Shunting_yard_algorithm

package parenthesis

import (
	"github.com/Goathy/containers/stack"
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
		ops     = stack.New[string]()
		postfix = make([]string, 0)
	)

	for _, token := range infix {
		switch token {
		case opLeftPar:
			ops.Push(token)
		case opRightPar:
			for {
				operator := ops.Pop()

				if operator == opLeftPar {
					break
				}

				postfix = append(postfix, operator)
			}
		case opAdd, opSub, opMulti, opDiv, opPow:
			for operator := ops.Peek(); !ops.IsEmpty() && precedence(operator) > precedence(token) || precedence(operator) == precedence(token) && assoc(token) == assocLeft; operator = ops.Peek() {
				operator = ops.Pop()
				postfix = append(postfix, operator)
			}
			ops.Push(token)
		default:
			postfix = append(postfix, token)
		}
	}

	for !ops.IsEmpty() {
		operator := ops.Pop()
		postfix = append(postfix, operator)
	}

	return postfix
}
