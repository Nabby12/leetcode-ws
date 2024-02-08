package main

import (
	"reflect"
	"testing"
)

func Test_trappingRainWater(t *testing.T) {
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
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			want: 6,
		},
		{
			name: "example2",
			input: input{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			want: 9,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := trappingRainWater(tt.input.height)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
