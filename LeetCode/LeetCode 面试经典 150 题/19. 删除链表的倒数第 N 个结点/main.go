package main

func main() {

}

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Val: -1, Next: head}
	pre := dummy
	p2 := head

	// p2向前走n步
	for i := 0; i < n; i++ {
		p2 = p2.Next
	}

	for p2 != nil {
		pre = pre.Next
		p2 = p2.Next
	}

	pnn := pre.Next.Next
	pre.Next = pnn

	return dummy.Next
}
