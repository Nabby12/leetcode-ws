package main

import "sort"

// https://leetcode.com/problems/car-fleet/description/
func carFleet4(target int, position []int, speed []int) int {
	cars := make([][2]int, len(position))
	for i, _ := range cars {
		cars[i] = [2]int{
			position[i],
			speed[i],
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})

	stack := []float64{}

	for _, car := range cars {
		goalTime := float64(target-car[0]) / float64(car[1])
		if len(stack) == 0 || stack[len(stack)-1] < goalTime {
			stack = append(stack, goalTime)
		}
	}

	return len(stack)
}
