package main

import (
	"reflect"
	"testing"
)

func Test_validAnagram20240127(t *testing.T) {
	t.Parallel()

	type input struct {
		s string
		t string
	}

	test := []struct {
		name  string
		input input
		want  bool
	}{
		{
			name: "example1",
			input: input{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "example2",
			input: input{
				s: "rat",
				t: "car",
			},
			want: false,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := validAnagram20240127(tt.input.s, tt.input.t)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
