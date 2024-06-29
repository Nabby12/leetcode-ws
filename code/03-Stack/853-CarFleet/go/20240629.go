package main

import (
	"sort"
)

// https://leetcode.com/problems/car-fleet/description/
func carFleet7(target int, position []int, speed []int) int {
	cars := [][2]int{}
	for i, p := range position {
		car := [2]int{}
		car[0] = p
		car[1] = speed[i]
		cars = append(cars, car)
	}
	sort.Slice(cars, func(i, j int) bool {
		return cars[i][0] > cars[j][0]
	})
	fleets := []float64{}
	for _, car := range cars {
		p, s := car[0], car[1]
		time := float64((target - p)) / float64(s)
		if len(fleets) > 0 && fleets[len(fleets)-1] >= time {
			continue
		}
		fleets = append(fleets, time)
	}
	return len(fleets)
}
