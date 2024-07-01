package main

import (
	"reflect"
	"testing"
)

func Test_twoSum2InputArrayIsSorted2(t *testing.T) {
	t.Parallel()

	type input struct {
		numbers []int
		target  int
	}

	test := []struct {
		name  string
		input input
		want  []int
	}{
		{
			name: "example1",
			input: input{
				numbers: []int{2, 7, 11, 15},
				target:  9,
			},
			want: []int{1, 2},
		},
		{
			name: "example2",
			input: input{
				numbers: []int{2, 3, 4},
				target:  6,
			},
			want: []int{1, 3},
		},
		{
			name: "example3",
			input: input{
				numbers: []int{-1, 0},
				target:  -1,
			},
			want: []int{1, 2},
		},
		{
			name: "example4",
			input: input{
				numbers: []int{5, 25, 75},
				target:  100,
			},
			want: []int{2, 3},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := twoSum2InputArrayIsSorted2(tt.input.numbers, tt.input.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
