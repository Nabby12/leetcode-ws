package main

import (
	"reflect"
	"testing"
)

func Test_searchInRotatedSortedArray(t *testing.T) {
	t.Parallel()

	type input struct {
		nums   []int
		target int
	}

	test := []struct {
		name  string
		input input
		want  int
	}{
		{
			name: "example1",
			input: input{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 0,
			},
			want: 4,
		},
		{
			name: "example2",
			input: input{
				nums:   []int{4, 5, 6, 7, 0, 1, 2},
				target: 3,
			},
			want: -1,
		},
		{
			name: "example3",
			input: input{
				nums:   []int{1},
				target: 0,
			},
			want: -1,
		},
		{
			name: "example4",
			input: input{
				nums:   []int{5, 1, 3},
				target: 3,
			},
			want: 2,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := searchInRotatedSortedArray(tt.input.nums, tt.input.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
