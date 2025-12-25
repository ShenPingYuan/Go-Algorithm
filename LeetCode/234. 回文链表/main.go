package main

import (
	"fmt"
)

func main() {
	// [1,0,0]
	head := &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0}}}
	fmt.Println(isPalindrome(head))
}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	if head.Next == nil {
		return true
	}

	if head.Next.Next == nil {
		if head.Next.Val == head.Val {
			return true
		}
		return false
	}

	dummy := &ListNode{Next: head}
	p1, p2 := dummy, dummy
	n := 0
	for p2 != nil && p2.Next != nil {
		p1 = p1.Next
		p2 = p2.Next.Next
		n++
	}
	dummy.Next = reverseN(head, n)

	if p2 == nil {
		p1 = p1.Next
	}
	p2 = head.Next

	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return true
}

func reverseN(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p1, p2 := head, head.Next
	head.Next = nil
	for n > 1 {
		pn := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = pn
		n--
	}
	head.Next = p2
	return p1
}

func isPalindrome2(head *ListNode) bool {
	left := head
	flag := true
	var travese func(head *ListNode)
	travese = func(head *ListNode) {
		if head == nil || flag == false {
			return
		}
		travese(head.Next)
		if head.Val != left.Val {
			flag = false
			return
		}
		left = left.Next
	}
	travese(head)
	return flag
}
