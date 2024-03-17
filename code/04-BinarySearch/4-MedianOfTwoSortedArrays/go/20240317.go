package main

import (
	"math"
)

// https://leetcode.com/problems/median-of-two-sorted-arrays/description/
func medianOfSortedArrays(nums1 []int, nums2 []int) float64 {
	var shorter, longer []int
	if len(nums1) <= len(nums2) {
		shorter, longer = nums1, nums2
	} else {
		longer, shorter = nums1, nums2
	}
	shorterLength, longerLength := len(shorter), len(longer)
	total := shorterLength + longerLength
	half := total / 2

	l, r := 0, shorterLength-1
	for {
		i := (l + r) / 2
		if l+r < 0 {
			i = -((l+r)/2 + 1)
		}
		j := half - i - 2

		var shorterLeft, shorterRight, longerLeft, longerRight int
		if i >= 0 {
			shorterLeft = shorter[i]
		} else {
			shorterLeft = -math.MaxInt
		}
		if i+1 < shorterLength {
			shorterRight = shorter[i+1]
		} else {
			shorterRight = math.MaxInt
		}
		if j >= 0 {
			longerLeft = longer[j]
		} else {
			longerLeft = -math.MaxInt
		}
		if j+1 < longerLength {
			longerRight = longer[j+1]
		} else {
			longerRight = math.MaxInt
		}

		if shorterLeft <= longerRight && longerLeft <= shorterRight {
			if total%2 == 0 {
				return float64(max(shorterLeft, longerLeft)+min(shorterRight, longerRight)) / 2
			}
			return float64(min(shorterRight, longerRight))
		} else if shorterLeft <= longerRight {
			l = i + 1
		} else {
			r = i - 1
		}
	}
}
