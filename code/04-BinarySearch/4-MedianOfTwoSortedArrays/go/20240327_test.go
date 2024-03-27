package main

import (
	"reflect"
	"testing"
)

func Test_medianOfSortedArrays8(t *testing.T) {
	t.Parallel()

	type input struct {
		nums1 []int
		nums2 []int
	}

	test := []struct {
		name  string
		input input
		want  float64
	}{
		{
			name: "example1",
			input: input{
				nums1: []int{1, 3},
				nums2: []int{2},
			},
			want: 2.00000,
		},
		{
			name: "example2",
			input: input{
				nums1: []int{1, 2},
				nums2: []int{3, 4},
			},
			want: 2.50000,
		},
		{
			name: "example3",
			input: input{
				nums1: []int{},
				nums2: []int{1},
			},
			want: 1.00000,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := medianOfSortedArrays8(tt.input.nums1, tt.input.nums2)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
