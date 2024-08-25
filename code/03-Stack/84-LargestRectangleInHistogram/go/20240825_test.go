package main

import (
	"reflect"
	"testing"
)

func Test_largestRectangleInHistogram12(t *testing.T) {
	t.Parallel()

	type input struct {
		heights []int
	}

	test := []struct {
		name  string
		input input
		want  int
	}{
		{
			name: "example1",
			input: input{
				heights: []int{2, 1, 5, 6, 2, 3},
			},
			want: 10,
		},
		{
			name: "example2",
			input: input{
				heights: []int{2, 4},
			},
			want: 4,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := largestRectangleInHistogram12(tt.input.heights)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
