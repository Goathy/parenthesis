package tokenizer

import (
	"fmt"

	"github.com/Goathy/containers/queue"
	"github.com/Goathy/parenthesis"
)

type queuer[V any] interface {
	Enqueue(v V)
	Dequeue() V
	IsEmpty() bool
	Peek() V
}

type tokenizer struct {
	merge  bool
	queue  queuer[string]
	output []string
}

func New() *tokenizer {
	return &tokenizer{
		merge:  false,
		queue:  queue.New[string](),
		output: make([]string, 0),
	}
}

func (t *tokenizer) Tokenize(expression string) []string {
	for _, exp := range expression {
		switch e := string(exp); e {
		case parenthesis.Blank:
			continue
		case parenthesis.OpLeftPar,
			parenthesis.OpRightPar,
			parenthesis.OpAdd,
			parenthesis.OpMulti,
			parenthesis.OpDiv,
			parenthesis.OpPow,
			parenthesis.OpSub:

			if t.queue.Peek() == parenthesis.Empty && e == parenthesis.OpAdd {
				continue
			}

			if t.queue.Peek() == parenthesis.OpSub && e == parenthesis.OpAdd {
				continue
			}

			if t.queue.Peek() == parenthesis.Empty && e == parenthesis.OpSub {
				t.queue.Enqueue(e)
				t.merge = true
				continue
			}

			if t.queue.Peek() == parenthesis.OpSub && e == parenthesis.OpSub {
				t.queue.Enqueue(e)
				t.merge = true
				continue
			}

			if t.queue.Peek() == parenthesis.OpLeftPar && e == parenthesis.OpSub {
				v := t.queue.Dequeue()
				t.move(v)
				t.queue.Enqueue(e)
				t.merge = true
				continue
			}

			if !t.isOperator(t.queue.Peek()) && t.queue.Peek() != parenthesis.Empty {
				v := t.queue.Dequeue()
				t.move(v)
			}

			if t.isOperator(t.queue.Peek()) {
				v := t.queue.Dequeue()
				t.move(v)
			}

			t.queue.Enqueue(e)
		default:
			if t.merge {
				v := t.queue.Dequeue()
				t.queue.Enqueue(fmt.Sprintf("%s%s", v, e))
				t.merge = false
				continue
			}

			if t.isOperator(t.queue.Peek()) {
				v := t.queue.Dequeue()
				t.move(v)
				t.queue.Enqueue(e)
				continue
			}

			if !t.isOperator(t.queue.Peek()) {
				v := t.queue.Dequeue()
				t.queue.Enqueue(fmt.Sprintf("%s%s", v, e))
				continue
			}

			t.queue.Enqueue(e)
		}
	}

	for !t.queue.IsEmpty() {
		v := t.queue.Dequeue()
		t.move(v)
	}

	return t.output
}

func (t *tokenizer) isOperator(o string) bool {
	switch o {
	case parenthesis.OpPow,
		parenthesis.OpMulti,
		parenthesis.OpDiv,
		parenthesis.OpAdd,
		parenthesis.OpSub,
		parenthesis.OpLeftPar,
		parenthesis.OpRightPar:
		return true
	default:
		return false
	}
}

func (t *tokenizer) move(v string) {
	t.output = append(t.output, v)
}
