package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	t.Parallel()

	type input struct {
		l1 *ListNode
		l2 *ListNode
	}

	test := []struct {
		name  string
		input input
		want  *ListNode
	}{
		{
			name: "example1",
			input: input{
				l1: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 3,
						},
					},
				},
				l2: &ListNode{
					Val: 5,
					Next: &ListNode{
						Val: 6,
						Next: &ListNode{
							Val: 4,
						},
					},
				},
			},
			want: &ListNode{
				Val: 7,
				Next: &ListNode{
					Val: 0,
					Next: &ListNode{
						Val: 8,
					},
				},
			},
		},
		{
			name: "example2",
			input: input{
				l1: &ListNode{
					Val: 0,
				},
				l2: &ListNode{
					Val: 0,
				},
			},
			want: &ListNode{
				Val: 0,
			},
		},
		{
			name: "example3",
			input: input{
				l1: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
								Next: &ListNode{
									Val: 9,
									Next: &ListNode{
										Val: 9,
										Next: &ListNode{
											Val: 9,
										},
									},
								},
							},
						},
					},
				},
				l2: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 9,
							},
						},
					},
				},
			},
			want: &ListNode{
				Val: 8,
				Next: &ListNode{
					Val: 9,
					Next: &ListNode{
						Val: 9,
						Next: &ListNode{
							Val: 9,
							Next: &ListNode{
								Val: 0,
								Next: &ListNode{
									Val: 0,
									Next: &ListNode{
										Val: 0,
										Next: &ListNode{
											Val: 1,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tt := range test {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			got := addTwoNumbers(tt.input.l1, tt.input.l2)
			if !reflect.DeepEqual(got, tt.want) {
				gotValues := fmt.Sprintf("%v", got.Val)
				wantValues := fmt.Sprintf("%v", tt.want.Val)
				for got.Next != nil {
					gotValues += ", "
					gotValues += fmt.Sprintf("%v", got.Next.Val)
					got = got.Next
				}
				for tt.want.Next != nil {
					wantValues += ", "
					wantValues += fmt.Sprintf("%v", tt.want.Next.Val)
					tt.want = tt.want.Next
				}
				t.Fatalf("\ngot : %v\nwant: %v", gotValues, wantValues)
			}
		})
	}
}
