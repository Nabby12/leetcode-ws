package main

import (
	"reflect"
	"strconv"
	"testing"
)

func Test_minstack(t *testing.T) {
	t.Parallel()

	type input struct {
		ops    []string
		inputs [][]int
	}

	test := []struct {
		name  string
		input input
		want  []string
	}{
		{
			name: "example1",
			input: input{
				ops:    []string{"MinStack", "push", "push", "push", "getMin", "pop", "top", "getMin"},
				inputs: [][]int{[]int{}, []int{-2}, []int{0}, []int{-3}, []int{}, []int{}, []int{}, []int{}},
			},
			want: []string{"null", "null", "null", "null", "-3", "null", "0", "-2"},
		},
		{
			name: "example2",
			input: input{
				ops:    []string{"MinStack", "push", "push", "push", "push", "getMin", "pop", "getMin", "pop", "getMin", "pop", "getMin"},
				inputs: [][]int{[]int{}, []int{2}, []int{0}, []int{3}, []int{0}, []int{}, []int{}, []int{}, []int{}, []int{}, []int{}, []int{}},
			},
			want: []string{"null", "null", "null", "null", "null", "0", "null", "0", "null", "0", "null", "2"},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var gots []string
			minStack := MinStack{}
			for i, op := range tt.input.ops {
				input := tt.input.inputs[i]
				ms, output := execOps(minStack, op, input)

				minStack = ms

				var got string
				if output == nil {
					got = "null"
				} else {
					got = strconv.Itoa(*output)
				}
				gots = append(gots, got)
			}
			if !reflect.DeepEqual(gots, tt.want) {
				t.Fatalf("\ngots: %v\nwant: %v", gots, tt.want)
			}
		})
	}
}

func execOps(minStack MinStack, op string, input []int) (MinStack, *int) {
	switch op {
	case "MinStack":
		return Constructor(), nil
	case "push":
		minStack.Push(input[0])
		return minStack, nil
	case "pop":
		minStack.Pop()
		return minStack, nil
	case "top":
		top := minStack.Top()
		return minStack, &top
	case "getMin":
		getMin := minStack.GetMin()
		return minStack, &getMin
	}
	return minStack, nil
}
