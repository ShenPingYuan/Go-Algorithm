package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	p1, p2 := head, head
	for p2 != nil && p2.Next != nil {
		p2 = p2.Next.Next
		p1 = p1.Next
		if p1 == p2 {
			p1 = head
			for p1 != p2 {
				p1 = p1.Next
				p2 = p2.Next
			}
			return p1
		}
	}
	return nil
}
