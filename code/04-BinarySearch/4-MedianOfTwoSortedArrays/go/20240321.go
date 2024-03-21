package main

import "math"

// https://leetcode.com/problems/median-of-two-sorted-arrays/description/
func medianOfSortedArrays3(nums1 []int, nums2 []int) float64 {
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
			sMid = -((-(l + r) + 1) / 2)
		}
		lMid := half - sMid - 1 - 1
		var sLeft, sRight, lLeft, lRight int
		if sMid >= 0 {
			sLeft = short[sMid]
		} else {
			sLeft = -math.MaxInt
		}
		if sMid+1 < sLen {
			sRight = short[sMid+1]
		} else {
			sRight = math.MaxInt
		}
		if lMid >= 0 {
			lLeft = long[lMid]
		} else {
			lLeft = -math.MaxInt
		}
		if lMid+1 < lLen {
			lRight = long[lMid+1]
		} else {
			lRight = math.MaxInt
		}
		if sLeft <= lRight && sRight >= lLeft {
			if total%2 == 0 {
				return float64(max(sLeft, lLeft)+min(sRight, lRight)) / 2
			}
			return float64(min(sRight, lRight))
		} else if sRight <= lLeft {
			l = sMid + 1
		} else {
			r = sMid - 1
		}
	}
}
