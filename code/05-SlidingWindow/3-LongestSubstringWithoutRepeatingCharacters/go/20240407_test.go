package main

import (
	"reflect"
	"testing"
)

func Test_longestSubstringWithoutRepeatingCharacters(t *testing.T) {
	t.Parallel()

	type input struct {
		s string
	}

	test := []struct {
		name  string
		input input
		want  int
	}{
		{
			name: "example1",
			input: input{
				s: "abcabcbb",
			},
			want: 3,
		},
		{
			name: "example2",
			input: input{
				s: "bbbbb",
			},
			want: 1,
		},
		{
			name: "example3",
			input: input{
				s: "pwwkew",
			},
			want: 3,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := longestSubstringWithoutRepeatingCharacters(tt.input.s)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
