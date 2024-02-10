package main

// https://leetcode.com/problems/min-stack/description/
type MinStack struct {
	data []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
}

func (this *MinStack) Pop() {
	this.data = this.data[0 : len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	minimum := this.data[0]
	for _, val := range this.data {
		minimum = min(val, minimum)
	}
	return minimum
}
