package main

import (
	"reflect"
	"testing"
)

func Test_dailyTemperatures7(t *testing.T) {
	t.Parallel()

	type input struct {
		temperatures []int
	}

	test := []struct {
		name  string
		input input
		want  []int
	}{
		{
			name: "example1",
			input: input{
				temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			},
			want: []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		{
			name: "example2",
			input: input{
				temperatures: []int{30, 40, 50, 60},
			},
			want: []int{1, 1, 1, 0},
		},
		{
			name: "example3",
			input: input{
				temperatures: []int{30, 60, 90},
			},
			want: []int{1, 1, 0},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := dailyTemperatures7(tt.input.temperatures)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
