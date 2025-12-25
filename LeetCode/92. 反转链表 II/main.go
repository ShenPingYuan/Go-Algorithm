package main

import "fmt"

func main() {
	// head = [1,2,3,4,5], left = 2, right = 4
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	left, right := 2, 4
	newHead := reverseBetween(head, left, right)
	for newHead != nil {
		fmt.Print(newHead.Val, " ")
		newHead = newHead.Next
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || left == right {
		return head
	}
	dummy := &ListNode{Next: head}
	p1 := dummy
	for i := 0; i < left-1; i++ {
		p1 = p1.Next
	}
	p1.Next = reverseN(p1.Next, right-left+1)
	return dummy.Next
}

func reverseN(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil || n <= 1 {
		return head
	}

	p1, p2 := head, head.Next
	for p1 != nil && n > 1 {
		pn := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = pn
		n--
	}
	head.Next = p2
	return p1
}
