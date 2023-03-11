package parenthesis

import (
	"fmt"

	"github.com/Goathy/containers/queue"
)

func isOperator(s string) bool {
	return s == opLeftPar || s == opRightPar || s == opAdd || s == opMulti || s == opDiv || s == opPow || s == opSub
}

// RESOURCE: https://stackoverflow.com/questions/46861254/infix-to-postfix-for-negative-numbers#46861656
func Tokenize(expression string) []string {
	var (
		merge  = false
		ns     = queue.New[string]()
		output = make([]string, 0)
	)

	for _, exp := range expression {
		switch e := string(exp); e {
		case blank:
			continue
		case opLeftPar, opRightPar, opAdd, opMulti, opDiv, opPow, opSub:

			if ns.Peek() == empty && e == opAdd {
				continue
			}

			if ns.Peek() == opSub && e == opAdd {
				continue
			}

			if ns.Peek() == empty && e == opSub {
				ns.Enqueue(e)
				merge = true
				continue
			}

			if !isOperator(ns.Peek()) && ns.Peek() != empty {
				v := ns.Dequeue()
				output = append(output, v)
			}

			if isOperator(ns.Peek()) {
				v := ns.Dequeue()
				output = append(output, v)
			}

			ns.Enqueue(e)
		default:

			if merge {
				v := ns.Dequeue()

				ns.Enqueue(fmt.Sprintf("%s%s", v, e))
				merge = false
				continue
			}

			if isOperator(ns.Peek()) {
				v := ns.Dequeue()
				output = append(output, v)
				ns.Enqueue(e)
			} else if !isOperator(ns.Peek()) {
				v := ns.Dequeue()
				ns.Enqueue(fmt.Sprintf("%s%s", v, e))
			} else {
				ns.Enqueue(e)
			}
		}
	}

	for !ns.IsEmpty() {
		v := ns.Dequeue()

		output = append(output, v)
	}

	return output
}
