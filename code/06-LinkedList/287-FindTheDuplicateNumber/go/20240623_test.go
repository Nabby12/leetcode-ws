package main

import (
	"reflect"
	"testing"
)

func Test_findTheDuplicateNumber2(t *testing.T) {
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
				nums: []int{1, 3, 4, 2, 2},
			},
			want: 2,
		},
		{
			name: "example2",
			input: input{
				nums: []int{3, 1, 3, 4, 2},
			},
			want: 3,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := findTheDuplicateNumber2(tt.input.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
