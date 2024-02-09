package main

import (
	"reflect"
	"testing"
)

func Test_validParentheses(t *testing.T) {
	t.Parallel()

	type input struct {
		s string
	}

	test := []struct {
		name  string
		input input
		want  bool
	}{
		{
			name: "example1",
			input: input{
				s: "()",
			},
			want: true,
		},
		{
			name: "example2",
			input: input{
				s: "()[]{}",
			},
			want: true,
		},
		{
			name: "example3",
			input: input{
				s: "(]",
			},
			want: false,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := validParentheses(tt.input.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
