package main

import (
	"reflect"
	"testing"
)

func Test_carFleet5(t *testing.T) {
	t.Parallel()

	type input struct {
		target   int
		position []int
		speed    []int
	}

	test := []struct {
		name  string
		input input
		want  int
	}{
		{
			name: "example1",
			input: input{
				target:   12,
				position: []int{10, 8, 0, 5, 3},
				speed:    []int{2, 4, 1, 1, 3},
			},
			want: 3,
		},
		{
			name: "example2",
			input: input{
				target:   10,
				position: []int{3},
				speed:    []int{3},
			},
			want: 1,
		},
		{
			name: "example3",
			input: input{
				target:   100,
				position: []int{0, 2, 4},
				speed:    []int{4, 2, 1},
			},
			want: 1,
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := carFleet5(tt.input.target, tt.input.position, tt.input.speed)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
