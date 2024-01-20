package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val: 3,
			},
		},
	}
	l2 := ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	ans := addTwoNumbers(&l1, &l2)
	fmt.Println(ans.Val)
	fmt.Println(ans.Next.Val)
	fmt.Println(ans.Next.Next.Val)
}

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
