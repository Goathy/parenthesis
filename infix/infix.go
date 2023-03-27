package infix

import (
	"fmt"
	"strings"

	"github.com/Goathy/containers/stack"
	"github.com/Goathy/parenthesis"
)

func Transform(postfix []string) string {
	var (
		s = stack.New[string]()
	)

	for _, token := range postfix {
		switch token {
		case parenthesis.OpAdd,
			parenthesis.OpMulti,
			parenthesis.OpDiv,
			parenthesis.OpPow,
			parenthesis.OpSub:
			b, a := s.Pop(), s.Pop()
			infix := fmt.Sprintf("(%s %s %s)", a, token, b)

			s.Push(infix)
		default:
			if strings.HasPrefix(token, parenthesis.OpSub) {
				infix := fmt.Sprintf("(%s)", token)
				s.Push(infix)
			} else {
				s.Push(token)
			}
		}
	}

	result := s.Pop()

	return result
}
