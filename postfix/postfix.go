// Implementation of Edsger Dijkstra's the "shunting yard" algorithm
// https://en.wikipedia.org/wiki/Shunting_yard_algorithm

package postfix

import (
	"github.com/Goathy/containers/stack"
	"github.com/Goathy/parenthesis"
)

func Transform(infix []string) []string {
	var (
		stack  = stack.New[string]()
		result = make([]string, 0)
	)
	for _, token := range infix {
		switch token {
		case parenthesis.OpLeftPar:
			stack.Push(token)
		case parenthesis.OpRightPar:
			for {
				o := stack.Pop()

				if o == parenthesis.OpLeftPar {
					break
				}

				result = append(result, o)
			}
		case parenthesis.OpAdd,
			parenthesis.OpSub,
			parenthesis.OpMulti,
			parenthesis.OpDiv,
			parenthesis.OpPow:
			for o := stack.Peek(); !stack.IsEmpty() && precedence(o) > precedence(token) || precedence(o) == precedence(token) && assoc(token) == parenthesis.AssocLeft; o = stack.Peek() {
				o = stack.Pop()
				result = append(result, o)

			}
			stack.Push(token)
		default:
			result = append(result, token)

		}
	}

	for !stack.IsEmpty() {
		operator := stack.Pop()
		result = append(result, operator)
	}

	return result
}

func assoc(o string) parenthesis.Associativity {
	switch o {
	case parenthesis.OpPow:
		return parenthesis.AssocRight
	default:
		return parenthesis.AssocLeft
	}
}

func precedence(op string) int {
	switch op {
	case parenthesis.OpPow:
		return 4
	case parenthesis.OpMulti, parenthesis.OpDiv:
		return 3
	case parenthesis.OpAdd, parenthesis.OpSub:
		return 2
	default:
		return 0
	}
}
