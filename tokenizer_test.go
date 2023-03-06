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
