package parenthesis

import (
	"strings"

	"github.com/Goathy/stack"
)

const (
	left  = 1
	right = -1
)

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

func associativity(op string) int {
	switch op {
	case "^":
		return right
	default:
		return left
	}
}

// TODO: Implement Edsger Dijkstra's the "shunting yard" algorithm
// https://en.wikipedia.org/wiki/Shunting_yard_algorithm
func Postfix(infix []string) string {
	var (
		s, _ = stack.New[string](-1)
		re   = strings.Builder{}
	)

	for _, in := range infix {
		switch in {
		case "(":
			s.Push(in)
		case ")":
			for {
				op, _ := s.Pop()

				if op == "(" {
					break
				}

				re.WriteString(op)
			}
		case "+", "-", "*", "/", "^":
			for op, _ := s.Peek(); !s.IsEmpty() && precedence(op) > precedence(in) || precedence(op) == precedence(in) && associativity(in) == left; op, _ = s.Peek() {
				op, _ = s.Pop()
				re.WriteString(op)
			}

			s.Push(in)
		default:
			re.WriteString(in)
		}
	}

	for !s.IsEmpty() {
		v, _ := s.Pop()
		re.WriteString(v)
	}

	return re.String()
}
