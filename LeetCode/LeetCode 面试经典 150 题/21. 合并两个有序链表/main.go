package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var result *ListNode = nil
	var current *ListNode = nil
	if l1.Val <= l2.Val {
		current = l1
		l1 = l1.Next
	} else {
		current = l2
		l2 = l2.Next
	}
	result = current

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			l1 = l1.Next
		} else {
			current.Next = l2
			l2 = l2.Next
		}
		current = current.Next
	}
	if l1 != nil {
		current.Next = l1
	}
	if l2 != nil {
		current.Next = l2
	}
	return result
}

func mergeTwoLists1(l1 *ListNode, l2 *ListNode) *ListNode {
	// 虚拟头结点
	dummy := &ListNode{-1, nil}
	p := dummy
	p1 := l1
	p2 := l2

	for p1 != nil && p2 != nil {

		// 比较 p1 和 p2 两个指针
		// 将值较小的的节点接到 p 指针
		if p1.Val > p2.Val {
			p.Next = p2
			p2 = p2.Next
		} else {
			p.Next = p1
			p1 = p1.Next
		}
		// p 指针不断前进
		p = p.Next
	}

	if p1 != nil {
		p.Next = p1
	}

	if p2 != nil {
		p.Next = p2
	}

	return dummy.Next
}
