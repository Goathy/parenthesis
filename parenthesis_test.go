package parenthesis_test

import (
	"errors"
	"reflect"
	"testing"

	"github.com/Goathy/parenthesis"
)

func TestValidation(t *testing.T) {
	var errTest = errors.New("validation error, unexpected parenthesis")

	tt := []struct {
		desc  string
		input string
		want  error
	}{
		{
			desc:  "case 1",
			input: "()",
			want:  nil,
		},
		{
			desc:  "case 2",
			input: "(",
			want:  errTest,
		},
		{
			desc:  "case 3",
			input: ")",
			want:  errTest,
		},
		{
			desc:  "case 4",
			input: "te(st)",
			want:  nil,
		},
		{
			desc:  "case 5",
			input: "tes)",
			want:  errTest,
		},
		{
			desc:  "case 6",
			input: "tes)()",
			want:  errTest,
		},
	}
	for _, tc := range tt {
		t.Run(tc.desc, func(t *testing.T) {
			err := parenthesis.Validate(tc.input)

			if !reflect.DeepEqual(err, tc.want) {
				t.Errorf("want %q, got %q", tc.want, err)
			}
		})
	}
}
