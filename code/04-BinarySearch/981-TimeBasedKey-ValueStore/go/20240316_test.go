package main

import (
	"fmt"
	"reflect"
	"testing"
)

type testInput struct {
	key       string
	value     string
	timestamp int
}

func Test_timeBasedKeyValueStore(t *testing.T) {
	t.Parallel()

	type input struct {
		ops    []string
		inputs []testInput
	}

	test := []struct {
		name  string
		input input
		want  []string
	}{
		{
			name: "example1",
			input: input{
				ops: []string{"TimeMap", "set", "get", "get", "set", "get", "get"},
				inputs: []testInput{
					{},
					{key: "foo", value: "bar", timestamp: 1},
					{key: "foo", timestamp: 1},
					{key: "foo", timestamp: 3},
					{key: "foo", value: "bar2", timestamp: 4},
					{key: "foo", timestamp: 4},
					{key: "foo", timestamp: 5},
				},
			},
			want: []string{"null", "null", "bar", "bar", "null", "bar2", "bar2"},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			var gots []string
			timeMap := TimeMap{}
			for i, op := range tt.input.ops {
				input := tt.input.inputs[i]
				// fmt.Println(op, input)
				tm, output := execOps(timeMap, op, input)

				timeMap = tm

				var got string
				if output == nil {
					got = "null"
				} else {
					got = *output
				}
				gots = append(gots, got)
			}
			if !reflect.DeepEqual(gots, tt.want) {
				t.Fatalf("\ngots: %v\nwant: %v", gots, tt.want)
			}
		})
	}
}

func execOps(timeMap TimeMap, op string, input testInput) (TimeMap, *string) {
	switch op {
	case "TimeMap":
		return Constructor(), nil
	case "set":
		timeMap.Set(input.key, input.value, input.timestamp)
		fmt.Println(timeMap)
		return timeMap, nil
	case "get":
		res := timeMap.Get(input.key, input.timestamp)
		return timeMap, &res
	}
	return timeMap, nil
}
