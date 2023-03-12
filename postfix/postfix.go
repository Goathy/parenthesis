// Implementation of Edsger Dijkstra's the "shunting yard" algorithm
// https://en.wikipedia.org/wiki/Shunting_yard_algorithm

package postfix

import (
	"github.com/Goathy/containers/stack"
	"github.com/Goathy/parenthesis"
)

func assoc(o string) parenthesis.Associativity {
	switch o {
	case "^":
		return parenthesis.AssocRight
	default:
		return parenthesis.AssocLeft
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

func Transform(infix []string) []string {
	var (
		ops     = stack.New[string]()
		postfix = make([]string, 0)
	)

	for _, token := range infix {
		switch token {
		case parenthesis.OpLeftPar:
			ops.Push(token)
		case parenthesis.OpRightPar:
			for {
				operator := ops.Pop()

				if operator == parenthesis.OpLeftPar {
					break
				}

				postfix = append(postfix, operator)
			}
		case parenthesis.OpAdd,
			parenthesis.OpSub,
			parenthesis.OpMulti,
			parenthesis.OpDiv,
			parenthesis.OpPow:
			for operator := ops.Peek(); !ops.IsEmpty() && precedence(operator) > precedence(token) || precedence(operator) == precedence(token) && assoc(token) == parenthesis.AssocLeft; operator = ops.Peek() {
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
