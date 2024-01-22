package main

import (
	"reflect"
	"testing"
)

func Test_containsDuplicate(t *testing.T) {
	t.Parallel()

	type input struct {
		nums []int
	}

	test := []struct {
		name  string
		input input
		want  bool
	}{
		{
			name: "example1",
			input: input{
				nums: []int{1, 2, 3, 1},
			},
			want: true,
		},
		{
			name: "example2",
			input: input{
				nums: []int{1, 2, 3, 4},
			},
			want: false,
		},
		{
			name: "example3",
			input: input{
				nums: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
			},
			want: true,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := containsDuplicate(tt.input.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
