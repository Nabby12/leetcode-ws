package main

import (
	"reflect"
	"testing"
)

func Test_threeSum7(t *testing.T) {
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
		{
			name: "example5",
			input: input{
				nums: []int{-2, 0, 1, 1, 2},
			},
			want: [][]int{
				{-2, 0, 2},
				{-2, 1, 1},
			},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := threeSum7(tt.input.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
