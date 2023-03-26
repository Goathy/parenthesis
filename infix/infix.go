package infix

import (
	"fmt"

	"github.com/Goathy/containers/stack"
	"github.com/Goathy/parenthesis"
)

func Transform(postfix []string) string {
	operatorStack := stack.New[string]()

	for _, token := range postfix {
		switch token {
		case parenthesis.OpAdd, parenthesis.OpMulti, parenthesis.OpDiv:
			b, a := operatorStack.Pop(), operatorStack.Pop()
			infix := fmt.Sprintf("(%s %s %s)", a, token, b)

			operatorStack.Push(infix)
		default:
			operatorStack.Push(token)
		}
	}

	result := operatorStack.Pop()

	return result
}
