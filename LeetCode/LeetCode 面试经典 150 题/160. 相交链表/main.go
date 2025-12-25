package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	m, n := 0, 0
	p1, p2 := headA, headB
	for p1 != nil {
		m++
		p1 = p1.Next
	}

	for p2 != nil {
		n++
		p2 = p2.Next
	}

	p1, p2 = headA, headB
	if m > n {
		for i := 0; i < m-n; i++ {
			p1 = p1.Next
		}
	}

	if n > m {
		for i := 0; i < n-m; i++ {
			p2 = p2.Next
		}
	}

	for p1 != nil {
		if p1 == p2 {
			return p1
		}
		p1 = p1.Next
		p2 = p2.Next
	}
	return nil
}

func getIntersectionNode0(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 == nil {
			p1 = headB
		} else {
			p1 = p1.Next
		}
		if p2 == nil {
			p2 = headA
		} else {
			p2 = p2.Next
		}
	}
	return p1
}

func getIntersectionNode1(headA, headB *ListNode) *ListNode {
	nodeMap := make(map[*ListNode]bool)
	p1, p2 := headA, headB
	for p1 != nil {
		nodeMap[p1] = true
		p1 = p1.Next
	}

	for p2 != nil {
		if nodeMap[p2] {
			return p2
		}
		p2 = p2.Next
	}

	return nil
}

func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != nil {
		for p2 != nil {
			if p1 == p2 {
				return p1
			}
			p2 = p2.Next
		}
		p1 = p1.Next
		p2 = headB
	}
	return nil
}
