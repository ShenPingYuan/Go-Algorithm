package main

import "fmt"

func main() {
	// head = [1,2,3,4,5],k=2
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	k := 2
	newHead := reverseKGroup2(head, k)
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

func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}

	nums := []*ListNode{}
	p := head
	for p != nil {
		pn := p.Next
		p.Next = nil
		nums = append(nums, p)
		p = pn
	}
	for pos := 0; pos+k < len(nums); pos += k {
		start := pos
		end := pos + k - 1
		for start < end {
			nums[start], nums[end] = nums[end], nums[start]
			start++
			end--
		}
	}
	dummy := &ListNode{}
	p = dummy
	for i := 0; i < len(nums); i++ {
		p.Next = nums[i]
		p = p.Next
	}
	return dummy.Next
}

func reverseKGroup3(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	a, b := head, head
	// 不足k个节点，直接返回
	for i := 0; i < k-1; i++ {
		b = b.Next
		if b == nil {
			return head
		}
	}
	newHead := reverseN(a, k)
	a.Next = reverseKGroup3(a.Next, k)
	return newHead
}

func reverseKGroup2(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 1 {
		return head
	}
	dummy := &ListNode{Next: head}
	pre := dummy
	for {
		pn := pre.Next
		p1, p2 := reverseNV(pre.Next, k)
		pre.Next = p1
		if p2 == nil {
			break
		}
		pre = pn
	}
	return dummy.Next
}

func reverseNV(head *ListNode, n int) (*ListNode, *ListNode) {
	if head == nil || head.Next == nil || n == 1 {
		return head, nil
	}
	// 如果链表长度不足n，就不进行反转
	p := head
	k := n
	for k > 1 {
		p = p.Next
		if p == nil {
			return head, nil
		}
		k--
	}

	p1, p2 := head, head.Next
	p1.Next = nil
	for n > 1 {
		pn := p2.Next
		p2.Next = p1
		p1 = p2
		p2 = pn
		n--
	}
	head.Next = p2
	return p1, p2
}

func reverseN(head *ListNode, n int) *ListNode {
	if head == nil || head.Next == nil || n == 1 {
		return head
	}
	p1, p2 := head, head.Next
	p1.Next = nil
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
