package main

import (
	"reflect"
	"testing"
)

func Test_evaluateReversePolishNotation4(t *testing.T) {
	t.Parallel()

	type input struct {
		tokens []string
	}

	test := []struct {
		name  string
		input input
		want  int
	}{
		{
			name: "example1",
			input: input{
				tokens: []string{"2", "1", "+", "3", "*"},
			},
			want: 9,
		},
		{
			name: "example2",
			input: input{
				tokens: []string{"4", "13", "5", "/", "+"},
			},
			want: 6,
		},
		{
			name: "example3",
			input: input{
				tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"},
			},
			want: 22,
		},
		{
			name: "example4",
			input: input{
				tokens: []string{"18"},
			},
			want: 18,
		}, {
			name: "example5",
			input: input{
				tokens: []string{"3", "11", "+", "5", "-"},
			},
			want: 9,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := evaluateReversePolishNotation4(tt.input.tokens)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
