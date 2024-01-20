package main

import (
	"reflect"
	"testing"
)

func Test_twoSum(t *testing.T) {
	t.Parallel()

	type input struct {
		nums   []int
		target int
	}

	test := []struct {
		name  string
		input input
		want  []int
	}{
		{
			name: "example1",
			input: input{
				nums:   []int{2, 7, 11, 15},
				target: 9,
			},
			want: []int{0, 1},
		},
		{
			name: "example2",
			input: input{
				nums:   []int{3, 2, 4},
				target: 6,
			},
			want: []int{1, 2},
		},
		{
			name: "example3",
			input: input{
				nums:   []int{3, 3},
				target: 6,
			},
			want: []int{0, 1},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := twoSum(tt.input.nums, tt.input.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
