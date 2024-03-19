package main

import "sort"

// https://leetcode.com/problems/car-fleet/description/
func carFleet5(target int, position []int, speed []int) int {
	var cars [][2]int
	for i, p := range position {
		car := [2]int{
			p,
			speed[i],
		}
		cars = append(cars, car)
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})

	stack := []float64{}
	for _, car := range cars {
		p, s := car[0], car[1]
		hour := float64(target-p) / float64(s)
		if len(stack) == 0 || stack[len(stack)-1] < hour {
			stack = append(stack, hour)
		}
	}
	return len(stack)
}
