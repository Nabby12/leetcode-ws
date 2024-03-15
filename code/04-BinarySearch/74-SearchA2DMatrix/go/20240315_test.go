package main

import (
	"reflect"
	"testing"
)

func Test_searchA2DMatrix6(t *testing.T) {
	t.Parallel()

	type input struct {
		matrix [][]int
		target int
	}

	test := []struct {
		name  string
		input input
		want  bool
	}{
		{
			name: "example1",
			input: input{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 3,
			},
			want: true,
		},
		{
			name: "example2",
			input: input{
				matrix: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
				target: 13,
			},
			want: false,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := searchA2DMatrix6(tt.input.matrix, tt.input.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
