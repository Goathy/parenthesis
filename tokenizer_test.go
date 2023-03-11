package parenthesis_test

import (
	"reflect"
	"testing"

	"github.com/Goathy/parenthesis"
)

func TestTokenizer(t *testing.T) {
	tt := []struct {
		desc  string
		input string
		want  []string
	}{
		{
			desc:  "53",
			input: "53",
			want:  []string{"53"},
		},
		{
			desc:  "(53)",
			input: "(53)",
			want:  []string{"(", "53", ")"},
		},
		{
			desc:  "2+2",
			input: "2+2",
			want:  []string{"2", "+", "2"},
		},
		{
			desc:  "4+8*3",
			input: "4+8*3",
			want:  []string{"4", "+", "8", "*", "3"},
		},
		{
			desc:  "(4+8)*3",
			input: "(4+8)*3",
			want:  []string{"(", "4", "+", "8", ")", "*", "3"},
		},
		{
			desc:  "6/2*(1+2)",
			input: "6/2*(1+2)",
			want:  []string{"6", "/", "2", "*", "(", "1", "+", "2", ")"},
		},
		{
			desc:  "8/2*(2+2)",
			input: "8/2*(2+2)",
			want:  []string{"8", "/", "2", "*", "(", "2", "+", "2", ")"},
		},
		{
			desc:  "6+9+4^2",
			input: "6+9+4^2",
			want:  []string{"6", "+", "9", "+", "4", "^", "2"},
		},
		{
			desc:  "5*(6^2-2)",
			input: "5*(6^2-2)",
			want:  []string{"5", "*", "(", "6", "^", "2", "-", "2", ")"},
		},
		{
			desc:  "4*8^2+11",
			input: "4*8^2+11",
			want:  []string{"4", "*", "8", "^", "2", "+", "11"},
		},
		{
			desc:  "46+(8*4)/2",
			input: "46+(8*4)/2",
			want:  []string{"46", "+", "(", "8", "*", "4", ")", "/", "2"},
		},
		{
			desc:  "6+9+(4*2+4^2)",
			input: "6+9+(4*2+4^2)",
			want:  []string{"6", "+", "9", "+", "(", "4", "*", "2", "+", "4", "^", "2", ")"},
		},
		{
			desc:  "7^2*(25+10/5)-13",
			input: "7^2*(25+10/5)-13",
			want:  []string{"7", "^", "2", "*", "(", "25", "+", "10", "/", "5", ")", "-", "13"},
		},
		{
			desc:  "10-7*(3+2)+7^2",
			input: "10-7*(3+2)+7^2",
			want:  []string{"10", "-", "7", "*", "(", "3", "+", "2", ")", "+", "7", "^", "2"},
		},
		// {
		// 	desc:  "5-3*(2^3-5+7*(-3))",
		// 	input: "5-3*(2^3-5+7*(-3))",
		// 	want:  []string{"5", "-", "3", "*", "(", "2", "^", "3", "-", "5", "+", "7", "*", "(", "-3", ")", ")"},
		// },
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := parenthesis.Tokenize(tc.input)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}

}
