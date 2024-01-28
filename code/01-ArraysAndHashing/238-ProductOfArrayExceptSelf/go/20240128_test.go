package main

import (
	"reflect"
	"testing"
)

func Test_productOfArrayExceptSelf(t *testing.T) {
	t.Parallel()

	type input struct {
		nums []int
	}

	test := []struct {
		name  string
		input input
		want  []int
	}{
		{
			name: "example1",
			input: input{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{24, 12, 8, 6},
		},
		{
			name: "example2",
			input: input{
				nums: []int{-1, 1, 0, -3, 3},
			},
			want: []int{0, 0, 9, 0, 0},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := productOfArrayExceptSelf(tt.input.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
