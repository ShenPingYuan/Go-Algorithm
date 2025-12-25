package main

import "fmt"

func main() {
	// [1,1,2]
	head := &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2}}}
	fmt.Println(deleteDuplicates(head))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	slow, fast := head, head.Next
	for fast != nil {
		if fast.Val != slow.Val {
			slow.Next = fast
			slow = slow.Next
		}
		fast = fast.Next
	}
	slow.Next = nil
	return head
}
