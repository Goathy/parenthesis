package parenthesis

import (
	"fmt"

	"github.com/Goathy/containers/queue"
)

func isOperator(s string) bool {
	return s == opLeftPar || s == opRightPar || s == opAdd || s == opMulti || s == opDiv || s == opPow || s == opSub
}

func Tokenize(expression string) []string {
	var (
		ns     = queue.New[string]()
		output = make([]string, 0)
	)

	for _, exp := range expression {
		switch e := string(exp); e {
		case opLeftPar, opRightPar, opAdd, opMulti, opDiv, opPow, opSub:
			if !isOperator(ns.Peek()) && ns.Peek() != "" {
				v := ns.Dequeue()
				output = append(output, v)
			}

			if isOperator(ns.Peek()) {
				v := ns.Dequeue()
				output = append(output, v)
			}

			// TODO: Remove type conversion
			ns.Enqueue(e)
		default:

			// TODO: If e is number and Peek is opSub concat

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
