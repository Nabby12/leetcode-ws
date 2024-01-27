package main

import (
	"reflect"
	"testing"
)

func Test_groupAnagram2(t *testing.T) {
	t.Parallel()

	type input struct {
		strs []string
	}

	test := []struct {
		name  string
		input input
		want  [][]string
	}{
		{
			name: "example1",
			input: input{
				strs: []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			},
			want: [][]string{{"bat"}, []string{"nat", "tan"}, []string{"ate", "eat", "tea"}},
		},
		{
			name: "example2",
			input: input{
				strs: []string{""},
			},
			want: [][]string{[]string{""}},
		},
		{
			name: "example3",
			input: input{
				strs: []string{"a"},
			},
			want: [][]string{[]string{"a"}},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := groupAnagram2(tt.input.strs)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("\ngot : %v\nwant: %v", got, tt.want)
			}
		})
	}
}
