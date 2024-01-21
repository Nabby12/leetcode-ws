package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := result
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := 0

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		if carry != 0 {
			sum += carry
		}

		carry = sum / 10
		if carry > 0 {
			sum = sum % 10
		}
		tmp.Next = &ListNode{Val: sum}
		tmp = tmp.Next
	}

	return result.Next
}
