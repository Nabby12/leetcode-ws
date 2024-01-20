package main

func twoSum(nums []int, target int) []int {
	if len(nums) == 2 {
		return []int{0, 1}
	}

	answers := []int{}
	numMap := map[int]int{}

	for i, num := range nums {
		if numMap[target-num] != 0 {
			answers = append(answers, numMap[target-num]-1)
			answers = append(answers, i)
			return answers
		}
		numMap[num] = i + 1
	}

	return answers
}
