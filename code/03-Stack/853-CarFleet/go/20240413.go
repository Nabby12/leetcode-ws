package main

import (
	"sort"
)

// https://leetcode.com/problems/car-fleet/description/
func carFleet6(target int, position []int, speed []int) int {
	cars := [][2]int{}
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
	timeStack := []float64{}
	for _, car := range cars {
		time := float64(target-car[0]) / float64(car[1])
		if len(timeStack) == 0 || float64(timeStack[len(timeStack)-1]) < time {
			timeStack = append(timeStack, time)
		}
	}
	return len(timeStack)
}
