package main

// https://leetcode.com/problems/longest-increasing-subsequence/description/
func longestIncreasingSubsequence(nums []int) int {
	tails := []int{nums[0]}
	for _, num := range nums[1:] {
		lastVal := tails[len(tails)-1]
		if lastVal < num {
			tails = append(tails, num)
			continue
		}

		idx := binarySearch(tails, num)

		if idx == len(tails) {
			tails = append(tails, num)
		} else {
			tails[idx] = num
		}
	}

	return len(tails)
}

func binarySearch(arr []int, num int) int {
	ans := -1

	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (l + r) / 2
		midVal := arr[mid]

		if num == midVal {
			ans = mid
			break
		} else if num < midVal {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}

	if ans == -1 {
		ans = l
	}

	return ans
}
