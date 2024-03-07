package main

import (
	"reflect"
	"testing"
)

func Test_kokoEatingBananas2(t *testing.T) {
	t.Parallel()

	type input struct {
		piles []int
		h     int
	}

	test := []struct {
		name  string
		input input
		want  int
	}{
		{
			name: "example1",
			input: input{
				piles: []int{3, 6, 7, 11},
				h:     8,
			},
			want: 4,
		},
		{
			name: "example2",
			input: input{
				piles: []int{30, 11, 23, 4, 20},
				h:     5,
			},
			want: 30,
		},
		{
			name: "example3",
			input: input{
				piles: []int{30, 11, 23, 4, 20},
				h:     6,
			},
			want: 23,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := kokoEatingBananas2(tt.input.piles, tt.input.h)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
