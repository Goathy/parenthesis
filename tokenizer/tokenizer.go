package tokenizer

import (
	"fmt"

	"github.com/Goathy/containers/queue"
	"github.com/Goathy/parenthesis"
)

func Tokenize(expression string) []string {
	var (
		merge  = false
		queue  = queue.New[string]()
		output = make([]string, 0)
	)

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

			if queue.Peek() == parenthesis.Empty && e == parenthesis.OpAdd {
				continue
			}

			if queue.Peek() == parenthesis.OpSub && e == parenthesis.OpAdd {
				continue
			}

			if queue.Peek() == parenthesis.Empty && e == parenthesis.OpSub {
				queue.Enqueue(e)
				merge = true
				continue
			}

			if queue.Peek() == parenthesis.OpSub && e == parenthesis.OpSub {
				queue.Enqueue(e)
				merge = true
				continue
			}

			if queue.Peek() == parenthesis.OpLeftPar && e == parenthesis.OpSub {
				v := queue.Dequeue()
				output = append(output, v)
				queue.Enqueue(e)
				merge = true
				continue
			}

			if !isOperator(queue.Peek()) && queue.Peek() != parenthesis.Empty {
				v := queue.Dequeue()
				output = append(output, v)
			}

			if isOperator(queue.Peek()) {
				v := queue.Dequeue()
				output = append(output, v)
			}

			queue.Enqueue(e)
		default:
			if merge {
				v := queue.Dequeue()
				queue.Enqueue(fmt.Sprintf("%s%s", v, e))
				merge = false
				continue
			}

			if isOperator(queue.Peek()) {
				v := queue.Dequeue()
				output = append(output, v)
				queue.Enqueue(e)
				continue
			}

			if !isOperator(queue.Peek()) {
				v := queue.Dequeue()
				queue.Enqueue(fmt.Sprintf("%s%s", v, e))
				continue
			}

			queue.Enqueue(e)
		}
	}

	for !queue.IsEmpty() {
		v := queue.Dequeue()
		output = append(output, v)
	}

	return output
}

func isOperator(o string) bool {
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
