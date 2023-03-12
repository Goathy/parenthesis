// Implementation of Edsger Dijkstra's the "shunting yard" algorithm
// https://en.wikipedia.org/wiki/Shunting_yard_algorithm

package postfix

import (
	"github.com/Goathy/containers/stack"
	"github.com/Goathy/parenthesis"
)

type stacker[V any] interface {
	IsEmpty() bool
	Push(v V)
	Pop() V
	Peek() V
}

type postfix struct {
	stack  stacker[string]
	output []string
}

func New() *postfix {
	return &postfix{
		stack:  stack.New[string](),
		output: make([]string, 0),
	}
}

func (p *postfix) Transform(infix []string) []string {
	for _, token := range infix {
		switch token {
		case parenthesis.OpLeftPar:
			p.stack.Push(token)
		case parenthesis.OpRightPar:
			for {
				o := p.stack.Pop()

				if o == parenthesis.OpLeftPar {
					break
				}

				p.move(o)
			}
		case parenthesis.OpAdd,
			parenthesis.OpSub,
			parenthesis.OpMulti,
			parenthesis.OpDiv,
			parenthesis.OpPow:
			for o := p.stack.Peek(); !p.stack.IsEmpty() && p.precedence(o) > p.precedence(token) || p.precedence(o) == p.precedence(token) && p.assoc(token) == parenthesis.AssocLeft; o = p.stack.Peek() {
				o = p.stack.Pop()
				p.move(o)
			}
			p.stack.Push(token)
		default:
			p.move(token)

		}
	}

	for !p.stack.IsEmpty() {
		operator := p.stack.Pop()
		p.move(operator)
	}

	return p.output
}

func (p *postfix) assoc(o string) parenthesis.Associativity {
	switch o {
	case parenthesis.OpPow:
		return parenthesis.AssocRight
	default:
		return parenthesis.AssocLeft
	}
}

func (p *postfix) precedence(op string) int {
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

func (p *postfix) move(v string) {
	p.output = append(p.output, v)
}
