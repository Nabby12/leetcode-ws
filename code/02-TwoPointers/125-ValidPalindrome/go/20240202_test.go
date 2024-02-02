package main

import (
	"reflect"
	"testing"
)

func Test_validPalindrome(t *testing.T) {
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
				s: "A man, a plan, a canal: Panama",
			},
			want: true,
		},
		{
			name: "example2",
			input: input{
				s: "race a car",
			},
			want: false,
		},
		{
			name: "example3",
			input: input{
				s: " ",
			},
			want: true,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := validPalindrome(tt.input.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
