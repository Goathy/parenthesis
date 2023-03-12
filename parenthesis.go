// Implementation of Edsger Dijkstra's the "shunting yard" algorithm
// https://en.wikipedia.org/wiki/Shunting_yard_algorithm

package parenthesis

import (
	"github.com/Goathy/containers/stack"
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
		case OpLeftPar:
			ops.Push(token)
		case OpRightPar:
			for {
				operator := ops.Pop()

				if operator == OpLeftPar {
					break
				}

				postfix = append(postfix, operator)
			}
		case OpAdd, OpSub, OpMulti, OpDiv, OpPow:
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
