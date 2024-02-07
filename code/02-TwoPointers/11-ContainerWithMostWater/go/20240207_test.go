package main

import (
	"reflect"
	"testing"
)

func Test_containerWithMostWater(t *testing.T) {
	t.Parallel()

	type input struct {
		height []int
	}

	test := []struct {
		name  string
		input input
		want  int
	}{
		{
			name: "example1",
			input: input{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
		{
			name: "example2",
			input: input{
				height: []int{1, 1},
			},
			want: 1,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := containerWithMostWater(tt.input.height)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
