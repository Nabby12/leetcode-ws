package main

import (
	"reflect"
	"testing"
)

func Test_longestIncreasingSubsequence(t *testing.T) {
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
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
		{
			name: "example2",
			input: input{
				nums: []int{0, 1, 0, 3, 2, 3},
			},
			want: 4,
		},
		{
			name: "example3",
			input: input{
				nums: []int{7, 7, 7, 7, 7, 7, 7},
			},
			want: 1,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := longestIncreasingSubsequence(tt.input.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
