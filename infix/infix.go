package infix

import (
	"fmt"
	"strings"

	"github.com/Goathy/containers/stack"
	"github.com/Goathy/parenthesis"
)

func Transform(postfix []string) string {
	operatorStack := stack.New[string]()

	for _, token := range postfix {
		switch token {
		case parenthesis.OpAdd, parenthesis.OpMulti, parenthesis.OpDiv, parenthesis.OpPow, parenthesis.OpSub:
			b, a := operatorStack.Pop(), operatorStack.Pop()
			infix := fmt.Sprintf("(%s %s %s)", a, token, b)

			operatorStack.Push(infix)
		default:
			if strings.HasPrefix(token, parenthesis.OpSub) {
				infix := fmt.Sprintf("(%s)", token)
				operatorStack.Push(infix)
			} else {
				operatorStack.Push(token)
			}
		}
	}

	result := operatorStack.Pop()

	return result
}
