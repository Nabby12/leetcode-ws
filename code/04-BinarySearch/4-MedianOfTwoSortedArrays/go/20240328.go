package main

import "math"

// https://leetcode.com/problems/median-of-two-sorted-arrays/submissions/1214925306/
func medianOfSortedArrays9(nums1 []int, nums2 []int) float64 {
	short, long := nums1, nums2
	sLen, lLen := len(short), len(long)
	if sLen > lLen {
		short, long = long, short
		sLen, lLen = lLen, sLen
	}
	total := sLen + lLen
	half := total / 2
	l, r := 0, sLen-1
	for {
		sMid := (l + r) / 2
		if (l + r) < 0 {
			sMid = -((-(l + r))/2 + 1)
		}
		lMid := half - sMid - 1 - 1
		var sLeft, sRight, lLeft, lRight int
		if sMid < 0 {
			sLeft = -math.MaxInt
		} else {
			sLeft = short[sMid]
		}
		if sMid+1 >= sLen {
			sRight = math.MaxInt
		} else {
			sRight = short[sMid+1]
		}
		if lMid < 0 {
			lLeft = -math.MaxInt
		} else {
			lLeft = long[lMid]
		}
		if lMid+1 >= lLen {
			lRight = math.MaxInt
		} else {
			lRight = long[lMid+1]
		}
		if sRight >= lLeft && lRight >= sLeft {
			if total%2 == 0 {
				return float64(max(sLeft, lLeft)+min(sRight, lRight)) / 2
			}
			return float64(min(sRight, lRight))
		} else if sLeft > lRight {
			r = sMid - 1
		} else {
			l = sMid + 1
		}
	}
}
