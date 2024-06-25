package main

import (
	"reflect"
	"testing"
)

func Test_bestTimeToBuyAndSellStock5(t *testing.T) {
	t.Parallel()

	type input struct {
		prices []int
	}

	test := []struct {
		name  string
		input input
		want  int
	}{
		{
			name: "example1",
			input: input{
				prices: []int{7, 1, 5, 3, 6, 4},
			},
			want: 5,
		},
		{
			name: "example2",
			input: input{
				prices: []int{7, 6, 4, 3, 1},
			},
			want: 0,
		},
		{
			name: "example3",
			input: input{
				prices: []int{1, 2},
			},
			want: 1,
		},
		{
			name: "example4",
			input: input{
				prices: []int{1, 2, 4, 2, 5, 7, 2, 4, 9, 0, 9},
			},
			want: 9,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := bestTimeToBuyAndSellStock5(tt.input.prices)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
