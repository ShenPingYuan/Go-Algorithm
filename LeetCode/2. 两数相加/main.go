package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	p := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		n := &ListNode{Val: carry}
		carry = 0

		if l1 != nil {
			n.Val += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			n.Val += l2.Val
			l2 = l2.Next
		}

		if n.Val >= 10 {
			carry = 1
			n.Val -= 10
		}
		p.Next = n
		p = p.Next
	}
	if carry == 1 {
		p.Next = &ListNode{Val: 1}
	}
	return dummy.Next
}
