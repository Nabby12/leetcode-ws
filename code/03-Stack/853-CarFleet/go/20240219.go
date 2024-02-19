package main

import (
	"sort"
)

// https://leetcode.com/problems/car-fleet/description/
func carFleet2(target int, position []int, speed []int) int {
	cars := make([][2]int, len(position))

	for i, p := range position {
		cars[i] = [2]int{
			p,
			speed[i],
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})

	stack := []float64{}
	for _, car := range cars {
		p, s := car[0], car[1]
		goal := float64(target-p) / float64(s)
		if len(stack) == 0 || goal > stack[len(stack)-1] {
			stack = append(stack, goal)
		}
	}

	return len(stack)
}
