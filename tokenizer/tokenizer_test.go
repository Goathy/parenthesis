package tokenizer_test

import (
	"reflect"
	"testing"

	"github.com/Goathy/parenthesis/tokenizer"
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
			desc:  "+7",
			input: "+7",
			want:  []string{"7"},
		},
		{
			desc:  "-7",
			input: "-7",
			want:  []string{"-7"},
		},
		{
			desc:  "3--7",
			input: "3--7",
			want:  []string{"3", "-", "-7"},
		},
		{
			desc:  "67+-9.3",
			input: "67+-9.3",
			want:  []string{"67", "+", "-", "9.3"},
		},
		{
			desc:  "5434-+23.677",
			input: "5434-+23.677",
			want:  []string{"5434", "-", "23.677"},
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
		{
			desc:  "5-3*(2^3-5+7*(-3))",
			input: "5-3*(2^3-5+7*(-3))",
			want:  []string{"5", "-", "3", "*", "(", "2", "^", "3", "-", "5", "+", "7", "*", "(", "-3", ")", ")"},
		},
		{
			desc:  "2*(1+(4*(2+1)+3))",
			input: "2*(1+(4*(2+1)+3))",
			want:  []string{"2", "*", "(", "1", "+", "(", "4", "*", "(", "2", "+", "1", ")", "+", "3", ")", ")"},
		},
		{
			desc:  "(3*5^2/15)-(5-2^2)",
			input: "(3*5^2/15)-(5-2^2)",
			want:  []string{"(", "3", "*", "5", "^", "2", "/", "15", ")", "-", "(", "5", "-", "2", "^", "2", ")"},
		},
		{
			desc:  "((3+2)^2+3)-9+3^2",
			input: "((3+2)^2+3)-9+3^2",
			want:  []string{"(", "(", "3", "+", "2", ")", "^", "2", "+", "3", ")", "-", "9", "+", "3", "^", "2"},
		},
		{
			desc:  "(18/3)^2+((13+7)*5^2)",
			input: "(18/3)^2+((13+7)*5^2)",
			want:  []string{"(", "18", "/", "3", ")", "^", "2", "+", "(", "(", "13", "+", "7", ")", "*", "5", "^", "2", ")"},
		},
		{
			desc:  "78+(30-0.5*(28+8))/6",
			input: "78+(30-0.5*(28+8))/6",
			want:  []string{"78", "+", "(", "30", "-", "0.5", "*", "(", "28", "+", "8", ")", ")", "/", "6"},
		},
		{
			desc:  "(5.9-5.3)*7.2+1.4^2",
			input: "(5.9-5.3)*7.2+1.4^2",
			want:  []string{"(", "5.9", "-", "5.3", ")", "*", "7.2", "+", "1.4", "^", "2"},
		},
		{
			desc:  "(2.1^2+5.2-7.2)*7.1",
			input: "(2.1^2+5.2-7.2)*7.1",
			want:  []string{"(", "2.1", "^", "2", "+", "5.2", "-", "7.2", ")", "*", "7.1"},
		},
		{
			desc:  "2*20/2+(3+4)*3^2-6+15",
			input: "2*20/2+(3+4)*3^2-6+15",
			want:  []string{"2", "*", "20", "/", "2", "+", "(", "3", "+", "4", ")", "*", "3", "^", "2", "-", "6", "+", "15"},
		},
		{
			desc:  "3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3",
			input: "3 + 4 * 2 / ( 1 - 5 ) ^ 2 ^ 3",
			want:  []string{"3", "+", "4", "*", "2", "/", "(", "1", "-", "5", ")", "^", "2", "^", "3"},
		},
	}

	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			got := tokenizer.Tokenize(tc.input)

			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("want %q, got %q", tc.want, got)
			}
		})
	}

}
