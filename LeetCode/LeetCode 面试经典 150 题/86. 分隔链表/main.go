package main

import (
	"fmt"
)

func main() {
	//[1,4,3,2,5,2]
	var head = &ListNode{}
	var p = head
	// var head = &ListNode{Val: 1, Next: nil}
	var nums = []int{1, 4, 3, 2, 5, 2}
	for _, v := range nums {
		node := &ListNode{Val: v, Next: nil}
		p.Next = node
		p = p.Next
	}
	partition(head.Next, 3)
	head = head.Next
	for head != nil {
		fmt.Println(head.Val)
		head = head.Next
	}
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	dummy1 := &ListNode{}
	l1 := dummy1
	dummy2 := &ListNode{}
	l2 := dummy2
	for head != nil {
		if head.Val < x {
			l1.Next = &ListNode{Val: head.Val}
			l1 = l1.Next
		} else {
			l2.Next = &ListNode{Val: head.Val}
			l2 = l2.Next
		}
		head = head.Next
	}
	l1.Next = dummy2.Next
	return dummy1.Next
}
