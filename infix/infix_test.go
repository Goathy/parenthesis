package infix_test

import (
	"reflect"
	"testing"

	"github.com/Goathy/parenthesis/infix"
)

func TestPostfix(t *testing.T) {
	tt := []struct {
		desc  string
		input []string
		want  string
	}{
		{
			desc:  "22+",
			input: []string{"2", "2", "+"},
			want:  "(2 + 2)",
		},
		{
			desc:  "483*+",
			input: []string{"4", "8", "3", "*", "+"},
			want:  "(4 + (8 * 3))",
		},
		{
			desc:  "6/2*(1+2)",
			input: []string{"6", "2", "/", "1", "2", "+", "*"},
			want:  "((6 / 2) * (1 + 2))",
		},
		{
			desc:  "8/2*(2+2)",
			input: []string{"8", "2", "/", "2", "2", "+", "*"},
			want:  "((8 / 2) * (2 + 2))",
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := infix.Transform(tc.input)

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}
}
