package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	fmt.Println(5 / 2)
	return removeByRecursive(head, val)
}

func removeByIndex(p *ListNode, index int) *ListNode {
	var next *ListNode
	current := p
	head := p
	var prev *ListNode
	// 1,2,6,3,4,5,6
	// 6,2,3,5,6,7
	for current != nil {
		next = current.Next
		if current.Val == index {
			if current == head {
				head = next
				current = next
			} else {
				prev.Next = next
				current = prev
			}
		}
		prev = current
		current = next

	}

	return head
}
func removeByRecursive(p *ListNode, index int) *ListNode {
	// 1,2,6,3,4,5,6
	// 6,2,3,5,6,7
	if p == nil {
		return p
	}

	p.Next = removeByRecursive(p.Next, index)
	if p.Val == index {
		return p.Next
	}
	return p
}
func Test_removeElements(t *testing.T) {
	type args struct {
		head *ListNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			// 1,2,6,3,4,5,6
			args: args{
				head: &ListNode{
					Val: 1,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val: 6,
										},
									},
								},
							},
						},
					},
				},
				val: 7,
			},
			want: &ListNode{
				Val: 1,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val: 3,
						Next: &ListNode{
							Val: 4,
							Next: &ListNode{
								Val: 5,
							},
						},
					},
				},
			},
		},
		{
			// 6,2,6,3,4,5,6
			args: args{
				head: &ListNode{
					Val: 6,
					Next: &ListNode{
						Val: 2,
						Next: &ListNode{
							Val: 6,
							Next: &ListNode{
								Val: 3,
								Next: &ListNode{
									Val: 4,
									Next: &ListNode{
										Val: 5,
										Next: &ListNode{
											Val: 6,
										},
									},
								},
							},
						},
					},
				},
				val: 6,
			},
			want: &ListNode{
				Val: 2,
				Next: &ListNode{
					Val: 3,
					Next: &ListNode{
						Val: 4,
						Next: &ListNode{
							Val: 5,
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeElements(tt.args.head, tt.args.val)
			assert.Equal(t, tt.want, got, tt.args)
		})
	}
}
