package main

import (
	"reflect"
	"testing"
)

func Test_generateParentheses(t *testing.T) {
	t.Parallel()

	type input struct {
		n int
	}

	test := []struct {
		name  string
		input input
		want  []string
	}{
		{
			name: "example1",
			input: input{
				n: 3,
			},
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		{
			name: "example2",
			input: input{
				n: 1,
			},
			want: []string{"()"},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := generateParentheses(tt.input.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
