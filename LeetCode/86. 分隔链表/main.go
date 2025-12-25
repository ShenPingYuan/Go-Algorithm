package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	dummy1 := &ListNode{}
	dummy2 := &ListNode{}
	p1 := dummy1
	p2 := dummy2

	for head != nil {
		if head.Val < x {
			p1.Next = head
			head = head.Next
			p1 = p1.Next
			p1.Next = nil
		} else {
			p2.Next = head
			head = head.Next
			p2 = p2.Next
			p2.Next = nil
		}
	}
	p1.Next = dummy2.Next
	return dummy1.Next
}
