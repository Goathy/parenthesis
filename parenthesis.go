package parenthesis

import (
	"errors"

	"github.com/Goathy/stack"
)

func Validate(in string) error {
	const (
		lpar = '('
		rpar = ')'
	)

	var errUnbalanced = errors.New("validation error, unbalanced  parenthesis")

	s, _ := stack.New[rune](-1)

	for _, r := range in {
		switch r {
		case lpar:
			s.Push(r)
		case rpar:
			par, err := s.Peek()
			if err == stack.EOS {
				return errUnbalanced
			}

			if par == rpar {
				break
			}
			s.Pop()
		}
	}

	if !s.IsEmpty() {
		return errUnbalanced
	}

	return nil
}
