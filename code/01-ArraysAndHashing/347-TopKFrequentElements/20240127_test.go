package main

import (
	"reflect"
	"testing"
)

func Test_topKFrequentElements(t *testing.T) {
	t.Parallel()

	type input struct {
		nums []int
		k    int
	}

	test := []struct {
		name  string
		input input
		want  []int
	}{
		{
			name: "example1",
			input: input{
				nums: []int{1, 1, 1, 2, 2, 3},
				k:    2,
			},
			want: []int{1, 2},
		},
		{
			name: "example2",
			input: input{
				nums: []int{1},
				k:    1,
			},
			want: []int{1},
		},
		{
			name: "example3",
			input: input{
				nums: []int{1, 2},
				k:    2,
			},
			want: []int{1, 2},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := topKFrequentElements(tt.input.nums, tt.input.k)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
