package main

import (
	"reflect"
	"testing"
)

func Test_findMinimumInRotatedSortedArray(t *testing.T) {
	t.Parallel()

	type input struct {
		nums []int
	}

	test := []struct {
		name  string
		input input
		want  int
	}{
		{
			name: "example1",
			input: input{
				nums: []int{3, 4, 5, 1, 2},
			},
			want: 1,
		},
		{
			name: "example2",
			input: input{
				nums: []int{4, 5, 6, 7, 0, 1, 2},
			},
			want: 0,
		},
		{
			name: "example3",
			input: input{
				nums: []int{11, 13, 15, 17},
			},
			want: 11,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := findMinimumInRotatedSortedArray(tt.input.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
