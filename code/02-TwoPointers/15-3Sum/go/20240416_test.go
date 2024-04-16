package main

import (
	"reflect"
	"testing"
)

func Test_threeSum6(t *testing.T) {
	t.Parallel()

	type input struct {
		nums []int
	}

	test := []struct {
		name  string
		input input
		want  [][]int
	}{
		{
			name: "example1",
			input: input{
				nums: []int{-1, 0, 1, 2, -1, -4},
			},
			want: [][]int{
				{-1, -1, 2},
				{-1, 0, 1},
			},
		},
		{
			name: "example2",
			input: input{
				nums: []int{0, 1, 1},
			},
			want: [][]int{},
		},
		{
			name: "example3",
			input: input{
				nums: []int{0, 0, 0},
			},
			want: [][]int{
				{0, 0, 0},
			},
		},
		{
			name: "example4",
			input: input{
				nums: []int{1, 2, -2, -1},
			},
			want: [][]int{},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := threeSum6(tt.input.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
