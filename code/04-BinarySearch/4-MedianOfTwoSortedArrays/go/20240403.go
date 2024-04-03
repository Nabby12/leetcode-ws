package main

import "math"

// https://leetcode.com/problems/median-of-two-sorted-arrays/description/
func medianOfSortedArrays11(nums1 []int, nums2 []int) float64 {
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
		var sL, sR, lL, lR int
		if sMid >= 0 {
			sL = short[sMid]
		} else {
			sL = -math.MaxInt
		}
		if sMid+1 < sLen {
			sR = short[sMid+1]
		} else {
			sR = math.MaxInt
		}
		if lMid >= 0 {
			lL = long[lMid]
		} else {
			lL = -math.MaxInt
		}
		if lMid+1 < lLen {
			lR = long[lMid+1]
		} else {
			lR = math.MaxInt
		}
		if sL <= lR && lL <= sR {
			if total%2 == 0 {
				return float64(max(sL, lL)+min(sR, lR)) / 2
			}
			return float64(min(sR, lR))
		} else if sL > lR {
			r = sMid - 1
		} else {
			l = sMid + 1
		}
	}
}
